package page

import (
	"fmt"
	"testing"
)

func TestScratch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Scratch() != nil {
		t.Errorf("Expected Scratch() to return nil, but got %v", p.Scratch())
	}
}
