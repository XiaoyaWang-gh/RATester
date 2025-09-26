package page

import (
	"fmt"
	"testing"
)

func TestPages_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	pages := nop.Pages()
	if pages != nil {
		t.Errorf("Expected nil, got %v", pages)
	}
}
