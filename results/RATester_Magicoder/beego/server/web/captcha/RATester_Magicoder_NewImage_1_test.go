package captcha

import (
	"fmt"
	"testing"
)

func TestNewImage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	digits := []byte{1, 2, 3, 4, 5}
	width := 800
	height := 600
	img := NewImage(digits, width, height)

	if img.Rect.Dx() != width {
		t.Errorf("Expected width %d, got %d", width, img.Rect.Dx())
	}

	if img.Rect.Dy() != height {
		t.Errorf("Expected height %d, got %d", height, img.Rect.Dy())
	}

	// Add more assertions as needed for other fields and methods of the Image struct
}
