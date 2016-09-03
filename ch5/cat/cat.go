// Write a program which mimics the Unix program cat.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func displayReader(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		l := s.Text()
		fmt.Println(l)
	}
	if err := s.Err(); err != nil {
		fmt.Println("error reading file")
	}
}

func main() {
	if len(os.Args) == 1 {
		displayReader(os.Stdin)
	}

	a := os.Args[1:]
	for _, v := range a {
		f, err := os.Open(v)
		if err != nil {
			fmt.Println("error reading", v)
			continue
		}
		displayReader(f)
	}
}
