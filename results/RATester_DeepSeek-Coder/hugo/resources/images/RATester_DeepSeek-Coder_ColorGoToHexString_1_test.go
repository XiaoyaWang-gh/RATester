package images

import (
	"fmt"
	"image/color"
	"testing"
)

func TestColorGoToHexString_1(t *testing.T) {
	t.Run("Test with fully opaque color", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		color := color.RGBA{R: 255, G: 0, B: 0, A: 255}
		expected := "#ff0000"
		result := ColorGoToHexString(color)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Test with fully transparent color", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		color := color.RGBA{R: 255, G: 0, B: 0, A: 0}
		expected := "#00000000"
		result := ColorGoToHexString(color)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Test with semi-transparent color", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		color := color.RGBA{R: 255, G: 0, B: 0, A: 128}
		expected := "#80000080"
		result := ColorGoToHexString(color)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})
}
