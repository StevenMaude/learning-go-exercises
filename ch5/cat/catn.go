// Write a program which mimics the Unix program cat.
// Make it support the -n flag, where each line is numbered.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func displayReader(r io.Reader, n bool) {
	s := bufio.NewScanner(r)
	for i := 1; s.Scan(); i++ {
		l := s.Text()
		if n {
			fmt.Println(i, l)
		} else {
			fmt.Println(l)
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println("error reading file")
	}
}

func main() {
	nFlag := flag.Bool("n", false, "display line count")
	flag.Parse()
	a := flag.Args()

	if len(a) == 0 {
		displayReader(os.Stdin, *nFlag)
	}

	for _, v := range a {
		f, err := os.Open(v)
		if err != nil {
			fmt.Println("error reading", v)
			continue
		}
		displayReader(f, *nFlag)
	}
}
