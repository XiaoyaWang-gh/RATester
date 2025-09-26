package pagination

import (
	"fmt"
	"testing"
)

func TestNums_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		nums: 100,
	}

	nums := paginator.Nums()

	if nums != 100 {
		t.Errorf("Expected 100, but got %d", nums)
	}
}
