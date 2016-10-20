// Write a function that calculates the average of a float64 slice.
package main

import "fmt"

func average(s []float64) float64 {
	var sum float64
	for _, v := range s {
		sum += v
	}
	avg := sum / float64(len(s))
	return avg
}

func main() {
	// Should give an average of 3.3225.
	s := []float64{1.23, 3.56, 2.5, 6}
	fmt.Println("Average of the values:", s, "is", average(s))
}
