package captcha

import (
	"fmt"
	"image"
	"testing"
)

func TestdrawDigit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		Paletted:  &image.Paletted{},
		numWidth:  10,
		numHeight: 10,
		dotSize:   10,
	}

	digit := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := 10
	y := 10

	img.drawDigit(digit, x, y)

	// Add assertions here to check if the method has been called correctly
}
