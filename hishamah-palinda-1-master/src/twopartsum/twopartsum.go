package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	// TODO sum a
	sum := 0
	for _, v := range a {
		sum += v
	}
	// TODO send result on res
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)
	sum1 := <- ch
	sum2 := <- ch

	// TODO Get the subtotals from the channel and return their sum
	return sum1 + sum2
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
