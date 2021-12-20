package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%f\n", z)
	}
	return z
}

func main() {
	fmt.Print(Sqrt(36))
}
