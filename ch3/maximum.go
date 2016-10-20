// Write a function that finds the maximum value in an int slice ([]int).
package main

import "fmt"

func findMax(is []int) int {
	max := is[0]
	for _, v := range is[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	s := []int{5, 13, 2, -6, 843}
	fmt.Println("Maximum of", s, "is", findMax(s))
}
