package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	var r, g, b, a uint8 = 0, 0, 0, 0
	b = uint8(x / 2)
	r = uint8(y / 2)
	a = b
	return color.RGBA{r, g, b, a}
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{255, 65}
	pic.ShowImage(m)
}
