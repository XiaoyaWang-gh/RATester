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
		t.Errorf("Expected a non-nil LayoutHandler, got nil")
	}

	if lh.cache == nil {
		t.Errorf("Expected a non-nil cache, got nil")
	}
}
