// Based on example in golang net package docs.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func echoLine(conn net.Conn) {
	// Is this defer OK here?
	defer conn.Close()
	s := bufio.NewScanner(conn)
	s.Scan()
	if err := s.Err(); err != nil && err != io.EOF {
		log.Println(err)
		return
	}
	fmt.Println(s.Text())
}

func main() {
	log.Println("Server running...")
	ln, err := net.Listen("tcp", ":8053")
	if err != nil {
		log.Fatal(err)
	}
	// Should this go after error?
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go echoLine(conn)
	}
}
