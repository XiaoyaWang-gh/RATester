package page

import (
	"fmt"
	"testing"
)

func TestMarkup_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	markup := nop.Markup()
	if markup != NopMarkup {
		t.Errorf("Expected Markup to return NopMarkup, but got %v", markup)
	}
}
