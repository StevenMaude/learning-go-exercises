package main

import "fmt"

func main() {
	items := []string{"a", "b", "a", "a", "a", "c", "d", "e", "f", "g"}
	previous, out := items[0], items[0]

	for _, item := range items[1:] {
		if item == previous {
			continue
		}
		out += " " + item
		previous = item
	}
	fmt.Println(out)
}
