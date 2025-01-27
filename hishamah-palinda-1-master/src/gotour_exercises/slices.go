package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	slice := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			slice[i] = append(slice[i], uint8((i+j)/2))
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
