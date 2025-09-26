package lazy

import (
	"fmt"
	"testing"
)

func TestNew_59(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	init := New()
	if init == nil {
		t.Error("Expected New() to return a non-nil Init, but it returned nil")
	}
}
