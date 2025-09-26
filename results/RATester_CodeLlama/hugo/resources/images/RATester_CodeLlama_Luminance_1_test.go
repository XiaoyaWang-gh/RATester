package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestLuminance_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Color{
		color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		hex:       "#ffffff",
		luminance: 1,
	}
	if c.Luminance() != 1 {
		t.Errorf("Luminance() = %v, want %v", c.Luminance(), 1)
	}
}
