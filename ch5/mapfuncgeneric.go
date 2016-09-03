// Use the answer from the earlier map exercise but now make it generic using interfaces.
// Make it at least work for ints and strings.

// Based on glancing at the solution.
package main

import "fmt"

type empty interface{}

func mapGeneric(f func(empty) empty, is []empty) {
	for i, v := range is {
		is[i] = f(v)
	}
}

func reverse(s string) string {
	r := make([]rune, len(s))
	for i, v := range s {
		r[len(s)-(1+i)] = v
	}
	return string(r)
}

func negate(y empty) empty {
	switch x := y.(type) {
	case int:
		return -x
	case string:
		s := reverse(x)
		return s
	}
	return y
}

func main() {
	is := []empty{1, 4, 9, 16, 25}
	fmt.Println("is before applying func:", is)
	mapGeneric(negate, is)
	fmt.Println("is after applying func:", is)
	strings := []empty{"foo", "bar", "baz"}
	fmt.Println("strings before applying func:", strings)
	mapGeneric(negate, strings)
	fmt.Println("strings after applying func:", strings)
}
