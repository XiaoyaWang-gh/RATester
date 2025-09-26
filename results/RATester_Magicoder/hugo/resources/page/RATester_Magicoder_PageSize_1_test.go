package page

import (
	"fmt"
	"testing"
)

func TestPageSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{
		size: 10,
	}

	expected := 10
	actual := paginator.PageSize()

	if actual != expected {
		t.Errorf("Expected PageSize to be %d, but got %d", expected, actual)
	}
}
