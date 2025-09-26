package webp

import (
	"fmt"
	"testing"
)

func TestSupports_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if Supports() {
		t.Error("Expected false, got true")
	}
}
