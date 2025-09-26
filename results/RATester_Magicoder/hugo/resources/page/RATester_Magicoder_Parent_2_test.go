package page

import (
	"fmt"
	"testing"
)

func TestParent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	parent := nop.Parent()
	if parent != nil {
		t.Errorf("Expected nil, got %v", parent)
	}
}
