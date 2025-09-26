package page

import (
	"fmt"
	"testing"
)

func TestPrev_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	prev := nop.Prev()
	if prev != nil {
		t.Errorf("Expected nil, got %v", prev)
	}
}
