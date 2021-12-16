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
func (i Image) At(w, h int) color.Color {
	return color.RGBA{55, 44, 245, 65}
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{255, 65}
	pic.ShowImage(m)
}
