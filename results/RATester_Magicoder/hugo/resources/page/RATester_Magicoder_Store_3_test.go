package page

import (
	"fmt"
	"testing"
)

func TestStore_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	scratch := nop.Store()
	if scratch != nil {
		t.Errorf("Expected nil, got %v", scratch)
	}
}
