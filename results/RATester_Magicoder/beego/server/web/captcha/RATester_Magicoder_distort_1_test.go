package captcha

import (
	"fmt"
	"image"
	"image/color"
	"testing"
)

func Testdistort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		Paletted: &image.Paletted{
			Pix:     []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			Stride:  10,
			Rect:    image.Rect(0, 0, 10, 10),
			Palette: color.Palette{color.RGBA{0, 0, 0, 0}, color.RGBA{1, 1, 1, 1}},
		},
		numWidth:  10,
		numHeight: 10,
		dotSize:   1,
	}

	amplude := float64(1)
	period := float64(1)

	img.distort(amplude, period)

	// Add assertions here to check the state of img after distort is called
}
