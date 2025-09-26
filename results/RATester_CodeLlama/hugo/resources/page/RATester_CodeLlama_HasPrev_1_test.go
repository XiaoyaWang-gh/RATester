package page

import (
	"fmt"
	"testing"
)

func TestHasPrev_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Pager{number: 1}
	if p.HasPrev() {
		t.Errorf("Pager with number 1 should not have a previous page")
	}
	p = &Pager{number: 2}
	if !p.HasPrev() {
		t.Errorf("Pager with number 2 should have a previous page")
	}
}
