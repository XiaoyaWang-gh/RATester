package layouts

import (
	"fmt"
	"testing"
)

func TestNewLayoutHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lh := NewLayoutHandler()

	if lh == nil {
		t.Error("Expected a LayoutHandler, got nil")
	}

	if lh.cache == nil {
		t.Error("Expected a cache, got nil")
	}
}
