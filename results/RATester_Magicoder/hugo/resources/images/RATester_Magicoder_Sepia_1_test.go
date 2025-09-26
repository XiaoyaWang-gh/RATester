package images

import (
	"fmt"
	"testing"
)

func TestSepia_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	filter := filters.Sepia(0.5)

	if filter == nil {
		t.Error("Expected filter to not be nil")
	}
}
