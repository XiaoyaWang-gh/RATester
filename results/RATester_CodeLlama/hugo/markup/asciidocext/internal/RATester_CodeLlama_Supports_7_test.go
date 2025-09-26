package internal

import (
	"fmt"
	"testing"
)

func TestSupports_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := &AsciidocConverter{}
	if a.Supports(nil) {
		t.Error("Expected false")
	}
}
