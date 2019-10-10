package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

// https://golang.org/pkg/image/#Image
// ColorModel returns the Image's color model.
// ColorModel() color.Model
// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
// Bounds() Rectangle
// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
// At(x, y int) color.Color

// Image struct with width and height
type Image struct {
	w, h int
}

// ColorModel returns color model
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns recetangle starting from 0,0
func (img Image) Bounds() image.Rectangle {
	// Use Rect to return Rectangle
	return image.Rect(0, 0, img.w, img.h)
}

// At returns color at specific point
func (img Image) At(x, y int) color.Color {
	imgFunc := func(x, y int) uint8 {
		return uint8((x * y) ^ 2/2)
	}
	v := imgFunc(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 64}
	pic.ShowImage(m)
}
