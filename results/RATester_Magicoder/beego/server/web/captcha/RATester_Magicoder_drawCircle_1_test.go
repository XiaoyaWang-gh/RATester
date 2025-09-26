package captcha

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func TestdrawCircle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		Paletted: &image.Paletted{
			Pix:     []uint8{},
			Stride:  0,
			Rect:    image.Rectangle{},
			Palette: color.Palette{},
		},
		numWidth:  0,
		numHeight: 0,
		dotSize:   0,
	}

	x := 0
	y := 0
	radius := 0
	colorIdx := uint8(0)

	img.drawCircle(x, y, radius, colorIdx)

	// Add assertions here
}
