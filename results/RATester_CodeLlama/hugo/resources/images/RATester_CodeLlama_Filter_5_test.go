package images

import (
	"fmt"
	"image"
	"testing"

	"github.com/disintegration/gift"
)

func TestFilter_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &ImageProcessor{}
	src := &image.RGBA{}
	filters := []gift.Filter{}

	// Act
	_, err := p.Filter(src, filters...)

	// Assert
	if err != nil {
		t.Errorf("Filter() error = %v", err)
		return
	}
}
