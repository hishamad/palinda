package main

import "fmt"

func fibonacci() func() int {
	fib := 1
	fib_1 := 0
	c := 0

	return func() int {
		if c == 0 {
			c++
			return fib_1
		} else if c == 1 {
			c++
			return fib
		} else {
			temp := fib
			fib += fib_1
			fib_1 = temp
			return fib
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
