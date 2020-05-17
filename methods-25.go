package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image 型
type Image struct{}

// At 実装
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 100, 100}
}

// Bounds 実装
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

// ColorModel 実装
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
