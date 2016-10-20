package main

import (
	"bufio"
	"fmt"
	"os"
	"stack"
	"strconv"
)

func main() {
	// Based on code in bufio docs.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter value or operator")
	s := stack.Stack{}

	for scanner.Scan() {
		input := scanner.Text()
		i, err := strconv.Atoi(input)

		if err != nil {
			secondOperand, err := s.Pop()
			if err != nil {
				fmt.Println("No value to pop")
				continue
			}

			firstOperand, err := s.Pop()
			if err != nil {
				fmt.Println("No value to pop")
				s.Push(secondOperand)
				continue
			}

			switch input {
			case "+":
				v := firstOperand + secondOperand
				fmt.Println("Result:", v)
				s.Push(v)
			case "-":
				v := firstOperand - secondOperand
				fmt.Println("Result:", v)
				s.Push(v)
			case "*":
				v := firstOperand * secondOperand
				fmt.Println("Result:", v)
				s.Push(v)
			case "/":
				if secondOperand == 0 {
					fmt.Println("Can't divide by 0")
					s.Push(firstOperand)
					s.Push(secondOperand)
					continue
				}
				// Stack only handles ints.
				v := int(firstOperand / secondOperand)
				fmt.Println("Result:", v)
				s.Push(v)

			default:
				s.Push(firstOperand)
				s.Push(secondOperand)
				fmt.Println("Bad input", input)
			}

		} else {
			err := s.Push(i)
			if err != nil {
				fmt.Println("Stack full")
				continue
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
