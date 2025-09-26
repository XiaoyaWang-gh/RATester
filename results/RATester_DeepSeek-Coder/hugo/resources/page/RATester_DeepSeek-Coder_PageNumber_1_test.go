package page

import (
	"fmt"
	"testing"
)

func TestPageNumber_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	paginator := &Paginator{}
	pager := &Pager{number: 5, Paginator: paginator}

	result := pager.PageNumber()
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
