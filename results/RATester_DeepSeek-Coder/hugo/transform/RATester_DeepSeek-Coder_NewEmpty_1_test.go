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

	t.Parallel()

	chain := NewEmpty()

	if len(chain) != 0 {
		t.Errorf("Expected length of chain to be 0, got %d", len(chain))
	}

	if cap(chain) != 20 {
		t.Errorf("Expected capacity of chain to be 20, got %d", cap(chain))
	}
}
