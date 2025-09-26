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

	extender := New()
	if extender == nil {
		t.Fatal("Expected a non-nil Extender, but got nil")
	}
}
