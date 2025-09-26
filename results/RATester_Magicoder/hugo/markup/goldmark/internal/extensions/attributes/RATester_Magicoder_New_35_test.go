package attributes

import (
	"fmt"
	"testing"
)

func TestNew_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	extender := New()
	if extender == nil {
		t.Error("Expected an Extender, but got nil")
	}
}
