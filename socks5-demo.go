package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

const (
	socksVersion     = 5
	layoutVer        = 0
	layoutNofMethods = 1
	layoutMethods    = 2
	layoutCommand    = 1
	layoutRSV        = 2
	layoutATYP       = 3
	layoutAddr       = 4
	typeConnect      = 1

	typeIPv4   = uint8(1)
	typeDomain = uint8(3)
	typeIPv6   = uint8(4)
	lenIPv4    = 3 + 1 + net.IPv4len + 2
	lenIPv6    = 3 + 1 + net.IPv6len + 2
	lenDmBase  = 3 + 1 + 1 + 2
)

type (
	Socks5Negotiation struct {
		Version      uint8
		NumOfMethods uint8
		Methods      []uint8
	}

	Socks5Request struct {
		Version         uint8
		Command         uint8
		RSV             uint8
		AddressType     uint8
		Address         string
		Port            uint16
		AddressWithPort string
		RawAddr         []byte
	}
)

func reply(conn net.Conn, bytes []byte) (err error) {
	if _, err := conn.Write(bytes); err != nil {
		log.Println("[RESP Error]", err)
	}
	return
}

// extract method negotiation header
// |VER | NMETHODS | METHODS  |
// +----+----------+----------+
// | 1  |    1     | 1 to 255 |
func extractNegotiation(conn net.Conn) (socks *Socks5Negotiation, err error) {
	buf := make([]byte, requestBuf)

	if _, err = io.ReadFull(conn, buf[:layoutNofMethods+1]); err != nil {
		return
	}

	version := uint8(buf[layoutVer])
	if version != 5 {
		err = errors.New("NOT a Socks5 request")
		return
	}
	nOfMethods := uint8(buf[layoutNofMethods])
	if nOfMethods == 0 {
		// do nothing
	} else if _, err = io.ReadFull(conn,
		buf[layoutMethods:layoutMethods+nOfMethods]); err != nil {
		return
	}

	socks = new(tnt.Socks5Negotiation)
	socks.Version = version
	socks.NumOfMethods = nOfMethods

	for i := uint8(0); i < nOfMethods; i++ {
		method := uint8(buf[layoutMethods+i])
		socks.Methods = append(socks.Methods, method)
	}

	return
}

// extract socks5 request
// +----+-----+-------+------+----------+----------+
// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
// +----+-----+-------+------+----------+----------+
func extractRequest(conn net.Conn) (socksReq *Socks5Request, err error) {
	buf := make([]byte, requestBuf)

	if _, err = io.ReadFull(conn, buf[:layoutATYP+1]); err != nil {
		return
	}

	version := uint8(buf[layoutVer])
	if version != 5 {
		err = errors.New("NOT a socks5 request")
		return
	}
	command := uint8(buf[layoutCommand])
	if command != typeConnect {
		err = errors.New("only CONNECT be able to accept")
		return
	}
	RSV := uint8(buf[layoutRSV])
	ATYP := uint8(buf[layoutATYP])
	var address string
	var addrEnd int
	var reqLen int

	switch ATYP {
	case typeIPv4:
		if _, err = io.ReadFull(conn,
			buf[layoutAddr:layoutAddr+lenIPv4]); err != nil {
			return
		}
		addrEnd = layoutAddr + lenIPv4
		address = net.IP(buf[layoutAddr:addrEnd]).String()
		reqLen = lenIPv4
	case typeIPv6:
		if _, err = io.ReadFull(conn,
			buf[layoutAddr:layoutAddr+lenIPv6]); err != nil {
			return
		}
		addrEnd = layoutAddr + lenIPv6
		address = net.IP(buf[layoutAddr:addrEnd]).String()
		reqLen = lenIPv6
	case typeDomain:
		if _, err = io.ReadFull(conn,
			buf[layoutAddr:layoutAddr+1]); err != nil {
			return
		}
		addrLen := int(buf[layoutAddr])
		addrEnd = layoutAddr + 1 + addrLen
		reqLen = addrLen + lenDmBase
		if _, err = io.ReadFull(conn,
			buf[layoutAddr+1:addrEnd]); err != nil {
			return
		}
		address = string(buf[layoutAddr+1 : addrEnd])
	default:
		err = fmt.Errorf("address type is Unknown: %d", ATYP)
		return
	}

	if _, err = io.ReadFull(conn, buf[addrEnd:addrEnd+2]); err != nil {
		return
	}

	socksReq = new(tnt.Socks5Request)
	socksReq.Version = version
	socksReq.Command = command
	socksReq.RSV = RSV
	socksReq.AddressType = ATYP
	socksReq.Address = address
	socksReq.Port = binary.BigEndian.Uint16(buf[addrEnd : addrEnd+2])
	socksReq.AddressWithPort = net.JoinHostPort(address,
		strconv.Itoa(int(socksReq.Port)))
	socksReq.RawAddr = buf[layoutATYP:reqLen]
	return
}

// reply the method selection
// |VER | METHOD |
// +----+--------+
// | 1  |   1    |
func replyNegotiation(conn net.Conn, socks *tnt.Socks5Negotiation) {
	// no authentication required
	reply(conn, []byte{socksVersion, 0x00})
}

// reply the request
// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
// +----+-----+-------+------+----------+----------+
// | 1  |  1  | X'00' |  1   | Variable |    2     |
func replyRequest(conn net.Conn, socksRequest *tnt.Socks5Request) {
	reply(conn, []byte{0x05, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x00, 0x80, 0x88})
}

// rountine of per connection
// https://www.ietf.org/rfc/rfc1928.txt
func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1. extract info about negotiation
	socks, err := extractNegotiation(conn)
	if err != nil {
		log.Println("[Negotiate Request Error]", err)
		return
	}
	log.Println(socks)

	// 2. confirm negotiation
	replyNegotiation(conn, socks)

	// 3. extract info about request
	socksRequest, err := extractRequest(conn)
	if err != nil {
		log.Println("[Extract Request Error]", err)
		return
	}
	log.Println(socksRequest)

	// 4. confirm the connection was established
	replyRequest(conn, socksRequest)

	// 5. pipe data ...
}