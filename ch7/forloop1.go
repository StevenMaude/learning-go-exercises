// Modify the program you created in exercise For-loop to use channels,
// in other words, the function called in the body should now be a goroutine
// and communication should happen via channels. You should not worry yourself
// on how the goroutine terminates.

package main

import "fmt"

func send(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	go send(c)
	for {
		i := <-c
		fmt.Println(i)
	}
}
