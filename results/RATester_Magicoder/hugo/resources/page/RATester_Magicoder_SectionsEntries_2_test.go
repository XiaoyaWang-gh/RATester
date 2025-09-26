package page

import (
	"fmt"
	"testing"
)

func TestSectionsEntries_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	result := p.SectionsEntries()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
