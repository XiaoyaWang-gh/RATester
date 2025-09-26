package render

import (
	"fmt"
	"testing"
)

func TestOrdinal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &hookBase{
		ordinal: 1,
	}

	if h.Ordinal() != 1 {
		t.Errorf("Expected Ordinal to be 1, but got %d", h.Ordinal())
	}
}
