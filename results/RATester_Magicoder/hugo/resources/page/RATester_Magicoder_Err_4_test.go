package page

import (
	"fmt"
	"testing"
)

func TestErr_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)

	if err := nop.Err(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
