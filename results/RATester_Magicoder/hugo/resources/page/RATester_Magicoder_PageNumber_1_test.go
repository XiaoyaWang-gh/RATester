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

	p := Pager{number: 1}
	if p.PageNumber() != 1 {
		t.Errorf("Expected 1, got %d", p.PageNumber())
	}
}
