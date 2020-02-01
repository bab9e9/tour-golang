package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	rows, cols int
}

func (m *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.rows, m.cols)
}

func (m *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(&m)
}

