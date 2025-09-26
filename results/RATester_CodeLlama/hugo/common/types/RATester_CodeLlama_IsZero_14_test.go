package types

import (
	"fmt"
	"testing"
)

func TestIsZero_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := LowHigh[string]{Low: 0, High: 0}
	if !l.IsZero() {
		t.Errorf("l.IsZero() = %v, want %v", l.IsZero(), true)
	}
}
