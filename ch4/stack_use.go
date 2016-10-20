// Porting code from Chapter 3 to work with the package version of stack.
package main

import (
	"fmt"
	"stack"
)

func main() {
	s := stack.Stack{}
	fmt.Println(s)
	s.Push(5)
	fmt.Println(s)
	s.Push(6)
	fmt.Println(s)
	s.Push(7)
	fmt.Println(s)
	s.Push(8)
	fmt.Println(s)
	s.Push(9)
	fmt.Println(s)
	err := s.Push(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	for i := 0; i < 6; i++ {
		fmt.Println(s.Pop())
	}
}
