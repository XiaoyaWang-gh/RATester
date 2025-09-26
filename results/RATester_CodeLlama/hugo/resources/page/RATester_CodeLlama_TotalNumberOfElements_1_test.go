package page

import (
	"fmt"
	"testing"
)

func TestTotalNumberOfElements_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Paginator{
		total: 10,
	}

	if p.TotalNumberOfElements() != 10 {
		t.Errorf("Expected 10, got %d", p.TotalNumberOfElements())
	}
}
