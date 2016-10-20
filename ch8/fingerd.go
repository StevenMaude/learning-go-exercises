package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"path"
)

func main() {
	log.Println("fingerd running")
	ln, err := net.Listen("tcp", ":79")
	if err != nil {
		log.Fatal(err)
	}

	errors := make(chan error)
	go logErrors(errors)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("fingerd errored:", err)
		}
		go handleConn(conn, errors)
	}
}

func logErrors(errors chan error) {
	for {
		select {
		case err := <-errors:
			log.Println(err)
		}
	}
}

func flushWriter(w *bufio.Writer, errors chan error) {
	if err := w.Flush(); err != nil {
		errors <- err
	}
}

func handleConn(conn net.Conn, errors chan error) {
	defer conn.Close()

	w := bufio.NewWriter(conn)
	defer flushWriter(w, errors)

	log.Println("Accepted connection")

	s := bufio.NewScanner(conn)
	s.Scan()
	if err := s.Err(); err != nil {
		errors <- err
		fmt.Fprintln(w, "Error reading input")
		return
	}
	username := s.Text()

	queryUser(username, w, errors)

}

func queryUser(username string, w *bufio.Writer, errors chan error) {
	user, err := user.Lookup(username)
	log.Println("Looked up", username)
	if err != nil {
		fmt.Fprintln(w, "User", username, "not found")
		return
	}

	path := path.Join(user.HomeDir, ".plan")
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(w, "No .plan found for", username)
		return
	}
	fmt.Fprintln(w, ".plan:")

	f := bufio.NewReader(file)
	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Fprintln(w, "Error reading .plan")
		errors <- err
	}
}
