package images

import (
	"fmt"
	"testing"
)

func TestProcess_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filters := &Filters{}
	filter := filters.Process("test")

	if filter == nil {
		t.Error("Expected a filter, but got nil")
	}
}
