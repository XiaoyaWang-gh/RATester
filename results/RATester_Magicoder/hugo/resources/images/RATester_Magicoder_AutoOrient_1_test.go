package images

import (
	"fmt"
	"testing"
)

func TestAutoOrient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	filter := filters.AutoOrient()

	if filter == nil {
		t.Error("Expected a filter, but got nil")
	}
}
