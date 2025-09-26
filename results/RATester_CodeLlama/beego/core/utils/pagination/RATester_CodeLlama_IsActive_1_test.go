package pagination

import (
	"fmt"
	"testing"
)

func TestIsActive_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{
		page: 1,
	}
	if !p.IsActive(1) {
		t.Errorf("IsActive(1) = false, want true")
	}
	if p.IsActive(2) {
		t.Errorf("IsActive(2) = true, want false")
	}
}
