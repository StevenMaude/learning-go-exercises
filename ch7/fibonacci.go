// Write a function that takes an int value and gives that many terms of the Fibonacci sequence.
// But now the twist: You must use channels.
package main

import "fmt"

func fib(i int, ch chan int) {
	x, y := 0, 1
	for j := 1; j <= i; j++ {
		ch <- y
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fib(10, ch)
	for i := range ch {
		fmt.Println(i)
	}
}
