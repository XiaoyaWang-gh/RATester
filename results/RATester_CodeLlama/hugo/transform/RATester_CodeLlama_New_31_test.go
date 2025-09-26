package transform

import (
	"fmt"
	"testing"
)

func TestNew_31(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	trs := []Transformer{
		func(ft FromTo) error {
			return nil
		},
		func(ft FromTo) error {
			return nil
		},
	}

	// Act
	chain := New(trs...)

	// Assert
	if len(chain) != len(trs) {
		t.Errorf("Expected %d, got %d", len(trs), len(chain))
	}
}
