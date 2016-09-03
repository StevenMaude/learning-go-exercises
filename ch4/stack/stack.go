// Create a proper package for your stack implementation, Push, Pop and the
// Stack type need to be exported.

// Create a simple stack which can hold a fixed number of ints.
// It does not have to grow beyond this limit.
// Define push -- put something on the stack -- and pop -- retrieve something from the stack -- functions.
// The stack should be a LIFO (last in, first out) stack.

// Write a String method which converts the stack to a string representation. The stack in the figure could be represented as: [0:m] [1:l] [2:k] .
package stack

import "fmt"

// Had to read the definition of stack and the opening line of a
// method definition to get started.
// It wasn't clear what was needed from the exercise (and it also
// uses a lot of things that aren't present).
type Stack struct {
	i    int
	vals [5]int
}

func (s *Stack) Push(k int) error {
	if s.i >= len(s.vals) {
		return fmt.Errorf("stack full")
	}
	s.vals[s.i] = k
	s.i++
	fmt.Println("pushed", s.vals[s.i-1], "now i is", s.i)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.i == 0 {
		return 0, fmt.Errorf("stack empty")
	}
	val := s.vals[s.i-1]
	s.i--
	fmt.Println("popped, now i is", s.i)
	return val, nil
}

func (s Stack) String() string {
	var out string
	for i := 0; i <= s.i-1; i++ {
		out += fmt.Sprintf("[%v:%v] ", i, s.vals[i])
	}
	return out
}

/*
func main() {
	s := Stack{}
	fmt.Println(s)
	s.push(5)
	fmt.Println(s)
	s.push(6)
	fmt.Println(s)
	s.push(7)
	fmt.Println(s)
	s.push(8)
	fmt.Println(s)
	s.push(9)
	fmt.Println(s)
	err := s.push(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	for i := 0; i < 6; i++ {
		fmt.Println(s.pop())
	}
}*/
