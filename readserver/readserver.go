package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
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
	// Set a deadline of 10s
	// When the dead line is over the connection closes itself
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	fmt.Println("called handle")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// We are using Fprintf to write back to the connection
		fmt.Fprintf(conn, "I heard you say : %s\n", line)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when its done?
	fmt.Println("Code got here")
}
