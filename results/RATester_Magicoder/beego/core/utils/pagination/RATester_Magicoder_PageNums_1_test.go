package pagination

import (
	"fmt"
	"testing"
)

func TestPageNums_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		PerPageNums: 10,
		nums:        100,
	}

	expected := 10
	actual := paginator.PageNums()

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
