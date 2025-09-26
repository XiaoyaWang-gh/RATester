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

	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	i := Image{}
	newImg := i.WithImage(img)

	if newImg.imageConfig == nil {
		t.Error("Expected imageConfig to be set, but it was nil")
	}

	if newImg.Spec != nil {
		t.Error("Expected Spec to be nil, but it was not")
	}

	if newImg.imageConfig.config.Width != 100 {
		t.Errorf("Expected width to be 100, but got %d", newImg.imageConfig.config.Width)
	}

	if newImg.imageConfig.config.Height != 100 {
		t.Errorf("Expected height to be 100, but got %d", newImg.imageConfig.config.Height)
	}
}
