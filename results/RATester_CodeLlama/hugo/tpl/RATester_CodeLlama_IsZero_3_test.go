package tpl

import (
	"fmt"
	"testing"
)

func TestIsZero_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var info ParseInfo
	if !info.IsZero() {
		t.Errorf("IsZero() = %v, want %v", info.IsZero(), true)
	}
}
