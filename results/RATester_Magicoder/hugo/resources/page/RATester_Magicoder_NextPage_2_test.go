package page

import (
	"fmt"
	"testing"
)

func TestNextPage_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	nextPage := nop.NextPage()
	if nextPage != nil {
		t.Errorf("Expected nil, got %v", nextPage)
	}
}
