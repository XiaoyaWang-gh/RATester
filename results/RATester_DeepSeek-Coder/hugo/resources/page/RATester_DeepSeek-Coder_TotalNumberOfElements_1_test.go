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

	paginator := &Paginator{
		total: 100,
	}

	total := paginator.TotalNumberOfElements()

	if total != 100 {
		t.Errorf("Expected 100, got %d", total)
	}
}
