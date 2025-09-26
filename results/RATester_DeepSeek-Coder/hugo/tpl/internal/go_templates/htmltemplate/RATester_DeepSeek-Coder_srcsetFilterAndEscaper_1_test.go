package template

import (
	"fmt"
	"testing"
)

func TestSrcsetFilterAndEscaper_1(t *testing.T) {
	t.Run("Test with contentTypeSrcset", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := srcsetFilterAndEscaper("image.jpg 1x", contentTypeSrcset)
		expected := "image.jpg 1x"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with contentTypeURL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := srcsetFilterAndEscaper("http://example.com/image.jpg 1x", contentTypeURL)
		expected := "http%3A%2F%2Fexample.com%2Fimage.jpg%201x"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Test with multiple srcset elements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := srcsetFilterAndEscaper("image1.jpg 1x, image2.jpg 2x, image3.jpg 3x", contentTypeSrcset)
		expected := "image1.jpg%201x,%20image2.jpg%202x,%20image3.jpg%203x"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
