package images

import (
	"fmt"
	"image"
	"image/draw"
	"testing"

	"github.com/disintegration/gift"
)

func TestDraw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var f textFilter
	var dst draw.Image
	var src image.Image
	var options *gift.Options

	// act
	f.Draw(dst, src, options)

	// assert
}
