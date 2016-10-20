// Rewrite the loop again so that it fills an array and then prints that
// array to the screen.
package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
}
