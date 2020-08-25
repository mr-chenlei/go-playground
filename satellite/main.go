package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"taas.com/atom/buf"

	isd "github.com/jbenet/go-is-domain"
)

var (
	local  = flag.String("local", "127.0.0.1:10070", "Local listening address.")
	galaxy = flag.String("galaxy", "127.0.0.1:10080", "Galaxy listening address.")
	target = flag.String("target", "", "Target address.")
)

var (
	logger     *log.Logger
	ss5ConnMap = make(map[string]*net.UDPConn, 0)
	locker     sync.RWMutex
)

const (
	ss5Deadline = 30 * time.Second
)

func lookup(addr string) (string, error) {
	var s []string
	var domain, port string
	if strings.Contains(addr, ":") {
		s = strings.Split(addr, ":")
		domain = s[0]
		port = s[1]
	} else {
		domain = addr
		port = "80"
	}
	if isd.IsDomain(domain) {
		mx, err := net.LookupIP(domain)
		if err != nil {
			return "", fmt.Errorf("failed to lookup domain %v, %v", domain, err)
		}
		return mx[0].String() + ":" + port, nil
	}
	return domain + ":" + port, nil
}

func handshakeGalaxy(galaxy, target string) (net.Conn, error) {
	conn, err := net.Dial("tcp", galaxy)
	if err != nil {
		return nil, fmt.Errorf("dial galaxy failed: %v", err)
	}
	if _, err := conn.Write([]byte{5, 2, 0, 1}); err != nil {
		return nil, fmt.Errorf("write authentication failed: %v", err)
	}
	reader := &buf.BufferedReader{Reader: buf.NewReader(conn)}
	buffer := buf.StackNew()
	if _, err := buffer.ReadFullFrom(reader, 2); err != nil {
		buffer.Release()
		return nil, fmt.Errorf("insufficient header")
	}
	h, p, err := net.SplitHostPort(target)
	if err != nil {
		return nil, err
	}
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid port range: %s", p)
	}
	request := []byte{5, 1, 0, 1}
	request = append(request, net.ParseIP(h).To4()...)
	request = append(request, []byte{byte(uint16(port) >> 8), byte((uint16(port) << 8) >> 8)}...)
	if _, err := conn.Write(request); err != nil {
		return nil, fmt.Errorf("write request failed: %v", err)
	}
	if _, err := buffer.ReadFullFrom(reader, 10); err != nil {
		buffer.Release()
		return nil, fmt.Errorf("insufficient response %v", err)
	}
	return conn, nil
}

func transportTCP(reader net.Conn, galaxy, target string) {
	galaxyConn, err := handshakeGalaxy(galaxy, target)
	if err != nil {
		logger.Println(err)
		return
	}
	go func() {
		defer reader.Close()
		defer galaxyConn.Close()

		if _, err := io.Copy(reader, galaxyConn); err != nil {
			logger.Printf("connection closed on source side, %v:%v -> %v:%v -> %v:%v, error: %v",
				reader.RemoteAddr().Network(), reader.RemoteAddr(),
				galaxyConn.LocalAddr().Network(), galaxyConn.LocalAddr(),
				galaxyConn.RemoteAddr().Network(), galaxyConn.RemoteAddr(),
				err)
		}
	}()
	go func() {
		defer reader.Close()
		defer galaxyConn.Close()

		if _, err := io.Copy(galaxyConn, reader); err != nil {
			logger.Printf("connection closed on galaxy side, %v:%v -> %v:%v -> %v:%v, error: %v",
				galaxyConn.RemoteAddr().Network(), galaxyConn.RemoteAddr(),
				galaxyConn.LocalAddr().Network(), galaxyConn.LocalAddr(),
				reader.RemoteAddr().Network(), reader.RemoteAddr(),
				err)
		}
	}()
}

func transportUDP(conn, ss5Conn *net.UDPConn, saddr *net.UDPAddr) {
	logger.Println("new udp connection", saddr.String(), "->", conn.LocalAddr(), "->", ss5Conn.LocalAddr(), "-> galaxy")

	b := make([]byte, 4096)
	for {
		if err := ss5Conn.SetDeadline(time.Now().Add(ss5Deadline)); err != nil {
			logger.Println("failed to update udp timeout,", saddr.String(), "->", conn.LocalAddr(), "->", ss5Conn.LocalAddr(), "-> galaxy")
		}

		n, _, err := ss5Conn.ReadFrom(b)
		if err != nil {
			logger.Println("failed receive from galaxy,", err, "local:", ss5Conn.LocalAddr())
			break
		}

		if _, err := conn.WriteTo(b[10:n], saddr); err != nil {
			logger.Printf("forward data from galaxy to %v failed, received data size from galaxy %v, %v", saddr.String(), n, err)
			break
		}
	}

	locker.Lock()
	delete(ss5ConnMap, saddr.String())
	ss5Conn.Close()
	locker.Unlock()

	logger.Println("udp connection released", saddr.String(), "->", conn.LocalAddr(), "->", ss5Conn.LocalAddr(), "-> galaxy")
}

func main() {
	flag.Parse()

	_ = os.Mkdir("./", 0755)
	logFile, err := os.OpenFile("./satellite.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if nil != err {
		panic(err)
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	laddr, err := net.ResolveTCPAddr("tcp", *local)
	if err != nil {
		logger.Panic(err)
	}
	ls, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		logger.Panic(err)
	}

	logger.Println("start working....")
	dest, err := lookup(*target)
	if err != nil {
		logger.Panic("failed ", err)
	}
	logger.Println("target", dest)

	// TCP part
	go func() {
		for {
			conn, err := ls.Accept()
			if err != nil {
				logger.Println(err)
			}
			logger.Println("accept tcp connection from", conn.RemoteAddr())

			go transportTCP(conn, *galaxy, dest)
		}
	}()
	// UDP part
	go func() {
		galaxyAddr, err := net.ResolveUDPAddr("udp4", *galaxy)
		if err != nil {
			logger.Panic("failed to resolve galaxy udp addr,", err)
		}
		targetAddr, err := net.ResolveUDPAddr("udp4", dest)
		if err != nil {
			logger.Panic("failed to resolve dest udp addr,", err)
		}
		laddr, err := net.ResolveUDPAddr("udp4", *local)
		if err != nil {
			logger.Panic("failed resolve local udp addr,", err)
		}
		conn, err := net.ListenUDP("udp4", laddr)
		if err != nil {
			logger.Panic("failed to listen upd,", err)
		}
		defer conn.Close()

		b := make([]byte, 4096)
		for {
			n, saddr, err := conn.ReadFromUDP(b)
			if err != nil {
				continue
			}

			go func() {
				ip := targetAddr.IP.To4()
				port := make([]byte, 2)
				binary.BigEndian.PutUint16(port, uint16(targetAddr.Port))

				length := 10 + n
				request := make([]byte, length)
				copy(request[0:4], []byte{0, 0, 0, 1})
				copy(request[4:8], ip)
				copy(request[8:10], port)
				copy(request[10:], b[:n])

				var ss5Conn *net.UDPConn
				locker.RLock()
				if c, ok := ss5ConnMap[saddr.String()]; !ok {
					locker.RUnlock()
					ss5Addr, err := net.ResolveUDPAddr("udp4", ":0")
					if err != nil {
						logger.Println("resolving socks5 udp socket failed,", err)
						return
					}
					ss5Conn, err = net.DialUDP("udp4", ss5Addr, galaxyAddr)
					if err != nil {
						logger.Println("failed to dial udp to galaxy,", err)
						return
					}
					locker.Lock()
					ss5ConnMap[saddr.String()] = ss5Conn
					locker.Unlock()

					go transportUDP(conn, ss5Conn, saddr)
				} else {
					locker.RUnlock()
					ss5Conn = c
				}

				if err := ss5Conn.SetDeadline(time.Now().Add(ss5Deadline)); err != nil {
					logger.Println("failed to update udp timeout,", saddr.String(), "->", conn.LocalAddr(), "->", ss5Conn.LocalAddr(), "-> galaxy")
				}
				if _, err := ss5Conn.Write(request[:length]); err != nil {
					logger.Println("write data with socks5 failed,", err)
				}
			}()
		}
	}()

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-osSignals
	}
}
