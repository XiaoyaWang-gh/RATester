package page

import (
	"fmt"
	"testing"
)

func TestTitle_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	title := nop.Title()
	if title != "" {
		t.Errorf("Expected Title to be empty, but got %s", title)
	}
}
