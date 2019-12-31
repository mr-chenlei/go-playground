package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Listen("unix", "./domain_socket")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	select {}
}
