package pagination

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

	p := &Paginator{}
	p.page = 1
	if p.HasPrev() {
		t.Error("HasPrev() should be false")
	}
	p.page = 2
	if !p.HasPrev() {
		t.Error("HasPrev() should be true")
	}
}
