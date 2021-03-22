// Stefan Nilsson 2013-02-27

// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
	"time"
	"fmt"
	"sync"
)

type ComplexFunc func(complex128) complex128

var Funcs []ComplexFunc = []ComplexFunc{
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

func main() {
	s := time.Now()
	wg:= new(sync.WaitGroup)
	wg.Add(len(Funcs))
	for n, fn := range Funcs {
		go makePic(fn,n, wg)
	}
	wg.Wait()
	duration := time.Since(s)
	fmt.Println(duration)
}
func makePic(fn ComplexFunc, n int, wg* sync.WaitGroup){
	err := CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024)
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	err = png.Encode(file, Julia(f, n))
	file.Close()
	return
}

// Julia returns an image of size n x n of the Julia set for f.
func util(f ComplexFunc, s float64, img image.Image, i int, minY int, maxY int, wg* sync.WaitGroup){
	for j := minY; j < maxY; j++ {
		n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
		r := uint8(0)
		g := uint8(0)
		b := uint8(n % 32 * 8)
		img.(*image.RGBA).Set(i, j, color.RGBA{r, g, b, 255})
	}
	wg.Done()
}

func Julia(f ComplexFunc, n int) image.Image {
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	s := float64(n / 4)
	wg := new(sync.WaitGroup)
	dY := bounds.Max.Y - bounds.Min.Y
	wg.Add(dY)
	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		go util(f, s, img, i, bounds.Min.Y, bounds.Max.Y, wg)
	}
	wg.Wait()
	return img
}
// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		z = f(z)
	}
	return
}
