package images

import (
	"fmt"
	"image"
	"testing"
)

func TestWithImage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var i Image
	var img image.Image

	// Act
	i.WithImage(img)

	// Assert
	// ...
}
