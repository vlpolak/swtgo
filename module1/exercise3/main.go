package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pc := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		pc[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			pc[i][j] = uint8((j + i) / 2)
		}
	}
	return pc
}

func main() {
	pic.Show(Pic)
}
