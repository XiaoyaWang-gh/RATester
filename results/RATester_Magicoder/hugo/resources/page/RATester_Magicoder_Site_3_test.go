package page

import (
	"fmt"
	"testing"
)

func TestSite_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	if nop.Site() != nil {
		t.Errorf("Expected nil, got %v", nop.Site())
	}
}
