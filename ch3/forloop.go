// Take what you did in (the previous) exercise to write the for loop and
// extend it a bit.
// Put the body of the for loop - the fmt.Printf - in a separate function.
package main

import "fmt"

func printLoop(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i <= 9; i++ {
		printLoop(i)
	}
}
