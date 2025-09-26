package page

import (
	"fmt"
	"testing"
)

func TestNext_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	next := nop.Next()
	if next != nil {
		t.Errorf("Expected nil, got %v", next)
	}
}
