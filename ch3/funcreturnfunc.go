package main

import "fmt"

// Write a function that returns a function that performs a +2 on integers.
func plusTwo() func(int) int {
	f := func(i int) int {
		return i + 2
	}
	return f
}

// Generalize the function from above and and create a plusX(x) which returns functions that add x to an integer.
func plusX(x int) func(int) int {
	f := func(i int) int {
		return x + i
	}
	return f
}

func main() {
	p := plusTwo()
	// Should be 4.
	fmt.Printf("%v\n", p(2))
	// Should be 5.
	t := plusX(3)
	fmt.Printf("%v\n", t(2))
}
