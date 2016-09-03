// Write a function that finds the maximum value in an int slice ([]int).
// Create a program that shows the maximum number and that works for both
// integers and floats.
package main

import "fmt"

type sI []int

type sF []float64

type Greater interface {
	Greater(int, int) bool
	Len() int
	Value(int) interface{}
}

func (s sI) Greater(i int, j int) bool {
	return s[i] > s[j]
}

func (s sI) Len() int {
	return len(s)
}

func (s sI) Value(i int) interface{} {
	return s[i]
}

func (s sF) Greater(i int, j int) bool {
	return s[i] > s[j]
}

func (s sF) Len() int {
	return len(s)
}

func (s sF) Value(i int) interface{} {
	return s[i]
}

func findMax(g Greater) interface{} {
	max_index := 0
	for i := 1; i < g.Len(); i++ {
		if g.Greater(i, max_index) {
			max_index = i
		}
	}
	return g.Value(max_index)
}

func main() {
	s := sI{850, 5, 843, -6, 7, 16, 842}
	fmt.Println("Maximum of", s, "is", findMax(s))
	f := sF{5.5, 13.223, 2.1, -6.4, 843.958}
	fmt.Println("Maximum of", f, "is", findMax(f))
}
