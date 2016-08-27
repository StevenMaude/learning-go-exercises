// Write a function that performs a bubble sort on a slice of ints.
package main

import "fmt"

func RecursiveBubblesort(s []int) []int {
	var swapped bool
	for i := 0; i <= len(s)-2; i++ {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			swapped = true
		}
	}
	if swapped {
		s = RecursiveBubblesort(s)
	}
	return s
}

// Doesn't need to return since not recursive; slice is sorted in place.
func IterativeBubblesort(s []int) {
	for {
		swapped := false
		for i := 0; i <= len(s)-2; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main() {
	s := []int{5, 3, 2, 8, 16, 1, 4}
	t := []int{5, 3, 2, 8, 16, 1, 4}
	u := RecursiveBubblesort(s)
	IterativeBubblesort(t)
	fmt.Println(t, u)
}
