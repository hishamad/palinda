package main

import (
	"fmt"
	"math"
)

const EPS = 1e-9

func Sqrt(x float64) float64 {
	z := x
	temp := -1.0
	for i:= 0; i < 10 && math.Abs(temp - z) > EPS; i++ {
		temp = z
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
