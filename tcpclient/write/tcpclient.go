package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()
	fmt.Fprintln(conn, "Hey I just dailed you ")
	// bs, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	log.Println(err)
	// }
	//
	// fmt.Println(string(bs))
}
