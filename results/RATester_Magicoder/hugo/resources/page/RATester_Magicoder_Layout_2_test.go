package page

import (
	"fmt"
	"testing"
)

func TestLayout_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Layout() != "" {
		t.Errorf("Expected Layout() to return an empty string, but got %s", p.Layout())
	}
}
