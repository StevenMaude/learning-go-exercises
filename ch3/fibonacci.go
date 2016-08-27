// Write a function that takes an int value and gives that many terms of the Fibonacci sequence.
package main

import "fmt"

func fib(i int) {
	x, y := 0, 1
	for j := 1; j <= i; j++ {
		fmt.Println(y)
		x, y = y, x+y
	}
}

func main() {
	fib(10)
}
