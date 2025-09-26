package captcha

import (
	"fmt"
	"image"
	"testing"
)

func TestfillWithCircles_1(t *testing.T) {
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

	n := 5
	maxradius := 10

	img.fillWithCircles(n, maxradius)

	// Add assertions here to check if the image was filled correctly
}
