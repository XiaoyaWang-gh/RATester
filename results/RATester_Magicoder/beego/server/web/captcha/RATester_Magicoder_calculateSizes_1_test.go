package captcha

import (
	"fmt"
	"image"
	"testing"
)

func TestcalculateSizes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	img := &Image{
		Paletted: &image.Paletted{},
	}

	tests := []struct {
		width  int
		height int
		ncount int
		want   int
	}{
		{100, 100, 1, 1},
		{200, 100, 2, 1},
		{100, 200, 3, 1},
		{200, 200, 4, 1},
	}

	for _, test := range tests {
		img.calculateSizes(test.width, test.height, test.ncount)
		if img.dotSize != test.want {
			t.Errorf("calculateSizes(%d, %d, %d) = %d, want %d", test.width, test.height, test.ncount, img.dotSize, test.want)
		}
	}
}
