// Write code to calculate the average of a float64 slice. In a later
// exercise you will make it into a function.
package main

import "fmt"

func main() {
	// Should give an average of 3.3225.
	s := []float64{1.23, 3.56, 2.5, 6}
	var sum float64
	for _, v := range s {
		sum += v
	}
	avg := sum / float64(len(s))
	fmt.Println("Average of the values:", s, "is", avg)
}
