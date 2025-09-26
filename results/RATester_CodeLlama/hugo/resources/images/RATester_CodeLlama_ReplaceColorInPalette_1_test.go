package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestReplaceColorInPalette_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var c color.Color
	var p color.Palette

	// Act
	ReplaceColorInPalette(c, p)

	// Assert
	// TODO: Add assertions.
}
