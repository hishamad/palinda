package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)
	// Error: The program didn't work because it's stuck waiting forever until the data could be recieved. This happens when there is no other goroutine that could recieve the data.
	// Solution: It's easily fixed by initaiting a new goroutine in order to pass the data through the channel and will be recivied because the program now is running concurrently
	go func(){
		ch <- "Hello world!"
	}()
	fmt.Println(<-ch)
}
