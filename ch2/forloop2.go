// Rewrite the loop from to use goto. The keyword for may not be used.
package main

import "fmt"

func main() {
	i := 0
Start:
	fmt.Println(i)
	i++
	if i <= 9 {
		goto Start
	}
}
