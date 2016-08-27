// Write a function that takes a variable number of ints and print each integer
// on a separate line.
package main

import "fmt"

func printVarInts(is ...int) {
	for _, v := range is {
		fmt.Println(v)
	}
}

func main() {
	printVarInts(5, 6, 7, 8)
}
