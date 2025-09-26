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
		t.Errorf("NewEmpty() = %v, want %v", len(chain), 0)
	}
}
