package types

import (
	"fmt"
	"testing"
)

func TestIsNil_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var v any
	if !IsNil(v) {
		t.Errorf("IsNil(%v) = false, want true", v)
	}
}
