package transform

import (
	"fmt"
	"testing"
)

func TestNewEmpty_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	chain := NewEmpty()

	if len(chain) != 0 {
		t.Errorf("Expected chain to be empty, but got length: %d", len(chain))
	}

	if cap(chain) != 20 {
		t.Errorf("Expected chain to have capacity 20, but got: %d", cap(chain))
	}
}
