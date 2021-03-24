package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.

func main(){ 
	wg := new(sync.WaitGroup)	
	ch := make(chan int)
	go Print(ch, wg)
	wg.Add(11)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	// Error: The program is terminated before it gets to print the last number, 11
	// Solution: We create a waitGroup that waits for the eleven operations to be done and only after that the program can be terminated
	wg.Wait()
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg* sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
		wg.Done()
	}
}
