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
/*
// This exemplifies the go lang technique where there are no classes, 
// but you can add "methods" to a type --- functions modifying an implicit argument.
// This is _only_ an exercise, but it bothers me that the Image type we are creating 
// doesn't have much "content". I.e., only rows, cols.
// The functions (methods) defined on the type provide all of the "data",
//
// But this means that you cannot set anything about the object except its row,cols 
// How would you go about transforming an image, for instance?
//
*/
