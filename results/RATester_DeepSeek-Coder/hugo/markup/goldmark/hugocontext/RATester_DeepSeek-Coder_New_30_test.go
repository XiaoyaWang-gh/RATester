package hugocontext

import (
	"fmt"
	"testing"
)

func TestNew_30(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ext := New()
	if ext == nil {
		t.Errorf("Expected a non-nil Extender, got nil")
	}
}
