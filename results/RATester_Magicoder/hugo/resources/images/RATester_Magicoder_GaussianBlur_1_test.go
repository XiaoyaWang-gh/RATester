package images

import (
	"fmt"
	"testing"
)

func TestGaussianBlur_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	sigma := 2.0
	filter := filters.GaussianBlur(sigma)

	if filter == nil {
		t.Error("Expected filter to not be nil")
	}

	// Add more specific tests here
}
