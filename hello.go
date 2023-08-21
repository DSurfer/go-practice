package main

import (
	"fmt"
	"runtime"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for x := 1.0; x < 10; x++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	os := runtime.GOOS
	fmt.Println(Sqrt(4))
	fmt.Println(os)
}
