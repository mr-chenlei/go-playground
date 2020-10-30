/*
A very simple TCP tun-client written in Go.
This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.
Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var remote = flag.String("remote", "0.0.0.0:12345", "The hostname or IP to connect to; defaults to \"localhost\".")

func main() {
	flag.Parse()

	fmt.Printf("Connecting to %s...\n", *remote)

	conn, err := net.Dial("tcp", *remote)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}

	go func() {
		// client send
	}()

	go func() {
		// client read
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("> ")
			text, _ := reader.ReadString('\n')

			conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
			_, err := conn.Write([]byte(text))
			if err != nil {
				fmt.Println("Error writing to stream.")
				break
			}
		}
	}()

	select {}
}
