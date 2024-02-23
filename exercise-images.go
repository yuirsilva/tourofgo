package main

import (
	"image"
	"image/color"
)

type Image struct{}

func (c Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (b Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{10, 10},
	}
}

func (a Image) At(x, y int) color.Color {
	v := uint8(x * y)

	return color.RGBA{v, v, 255, 255}
}

func main() {
	// m := Image{}
	// pic.ShowImage(m)
}

// https://go.dev/tour/methods/25
