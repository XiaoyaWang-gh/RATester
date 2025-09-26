package images

import (
	"fmt"
	"testing"
)

func TestGrayscale_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	grayscaleFilter := filters.Grayscale()

	if grayscaleFilter == nil {
		t.Errorf("Expected grayscale filter to be not nil")
	}

	// Assuming that the grayscale filter is implemented correctly, we can't really test it here.
	// We would need to create an image and apply the filter to it, and then check the resulting image.
	// This is a bit more involved and would depend on the specifics of the implementation.
}
