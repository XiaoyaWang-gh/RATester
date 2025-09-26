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
		t.Error("Expected HasPrev() to return false, but it returned true")
	}

	p = &Pager{number: 2}
	if !p.HasPrev() {
		t.Error("Expected HasPrev() to return true, but it returned false")
	}
}
