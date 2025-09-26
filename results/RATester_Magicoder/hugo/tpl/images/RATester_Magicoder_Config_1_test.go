package images

import (
	"fmt"
	"image"
	"testing"

	"github.com/spf13/afero"
)

func TestConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{
		readFileFs: afero.NewMemMapFs(),
		cache:      make(map[string]image.Config),
	}

	// Test with a valid file
	afero.WriteFile(ns.readFileFs, "test.jpg", []byte("image data"), 0644)
	config, err := ns.Config("test.jpg")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if config.Width != 1 || config.Height != 1 {
		t.Errorf("Expected config width and height to be 1, got %d, %d", config.Width, config.Height)
	}

	// Test with an invalid file
	_, err = ns.Config("nonexistent.jpg")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test with an empty filename
	_, err = ns.Config("")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test with a non-string path
	_, err = ns.Config(123)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
