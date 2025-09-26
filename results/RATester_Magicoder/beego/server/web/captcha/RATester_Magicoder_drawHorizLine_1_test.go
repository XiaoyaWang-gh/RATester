package captcha

import (
	"fmt"
	"image"
	"testing"
)

func TestdrawHorizLine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		Paletted:  &image.Paletted{},
		numWidth:  10,
		numHeight: 10,
		dotSize:   1,
	}

	img.drawHorizLine(1, 5, 1, 1)

	for x := 1; x <= 5; x++ {
		if img.ColorIndexAt(x, 1) != 1 {
			t.Errorf("Expected color index 1 at (%d, 1), but got %d", x, img.ColorIndexAt(x, 1))
		}
	}
}
