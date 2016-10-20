// Write a simple map()-function in Go.
// It is sufficient for this function only to work for ints.
package main

import "fmt"

func mapInts(f func(int) int, is []int) {
	for i, v := range is {
		is[i] = f(v)
	}
}

func negate(i int) int {
	return -i
}

func main() {
	is := []int{1, 4, 9, 16, 25}
	fmt.Println("is before applying func:", is)
	mapInts(negate, is)
	fmt.Println("is after applying func:", is)
}
