package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}
		_, _ = io.WriteString(conn, "Hello from tcp server, How are you\n")
		_, _ = fmt.Fprintf(conn, "Well I hope")
		conn.Close()
	}
}
