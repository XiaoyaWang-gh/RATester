package page

import (
	"fmt"
	"testing"
)

func TestFirstSection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	result := nop.FirstSection()
	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
