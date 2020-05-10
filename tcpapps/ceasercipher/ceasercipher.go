package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Panic(err)
		}
		go handle(conn)

	}
}

func handle(conn net.Conn) {

	fmt.Println("called handle")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		r := rot13(lineBytes)

		fmt.Fprintf(conn, "%s - %s\n\n", lineBytes, r)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when its done?
	fmt.Println("Code got here")
}
func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return r13
}
