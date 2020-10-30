package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
)

func main() {
	list := NewIndexBuffer(32*1024, 2048)
	for _, v := range list {
		fmt.Println(v.index, v.crc32, v.buffer)
	}
}

const (
	indexLen = 4
	crc32Len = 4
)

type CRCBuffer struct {
	index  uint32
	crc32  uint32
	buffer []byte
}

func CRC32Verify(buffer []byte) uint32 {
	ieee := crc32.NewIEEE()
	_, err := io.WriteString(ieee, string(buffer))
	if err != nil {
		panic(err)
	}
	return ieee.Sum32()
}

func NewIndexBuffer(size, split int) []*CRCBuffer {
	buf := newBuffer(size - split*(indexLen+crc32Len))
	return splitBuffer(buf, split)
}

func splitBuffer(buffer []byte, split int) []*CRCBuffer {
	list := make([]*CRCBuffer, 0)
	bufLen := len(buffer) / split
	var index uint32
	for i := 1; i < split; i++ {
		crcBuffer := &CRCBuffer{}
		// Buffer
		tmp := buffer[(i-1)*bufLen : i*bufLen]

		// Index
		crcBuffer.index = index
		// Index to []byte
		indexByte := make([]byte, 4)
		binary.BigEndian.PutUint32(indexByte, index)
		// CRC32
		crcBuffer.crc32 = CRC32Verify(tmp)
		// CRC32 to []byte
		crc32Byte := make([]byte, 4)
		binary.BigEndian.PutUint32(crc32Byte, crcBuffer.crc32)

		// Append Index, CRC32, tmp to crcBuffer.buffer
		crcBuffer.buffer = append(crcBuffer.buffer, indexByte...)
		crcBuffer.buffer = append(crcBuffer.buffer, crc32Byte...)
		crcBuffer.buffer = append(crcBuffer.buffer, tmp...)

		list = append(list, crcBuffer)

		index++
	}
	return list
}

func newBuffer(length int) []byte {
	if length == 0 {
		return nil
	}
	var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+`~1234567890-====")
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for newBuffer()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip this number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return b
			}
		}
	}
}
