// There are a few annoying issues left if you resolve question 1 above.
// One of the problems is that the goroutine isn't neatly cleaned up when
// main.main() exits. And worse, due to a race condition between the exit of
// main.main() and main.shower() not all numbers are printed. It should print
// up until 9, but sometimes it prints only to 8. Adding a second quit-channel
// you can remedy both issues. Do this.

// NB: I think an easier solution for the code I wrote is to just close the
// channel in the goroutine and either iterate over the channel values via
// range, or just check if the value is ok and break if not.
package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan int)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- 0
}

func shower(ch chan int, quit chan int) {
	for {
		select {
		case j := <-ch:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}
