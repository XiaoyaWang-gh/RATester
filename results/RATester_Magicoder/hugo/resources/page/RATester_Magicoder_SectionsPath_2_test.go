package page

import (
	"fmt"
	"testing"
)

func TestSectionsPath_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.SectionsPath() != "" {
		t.Errorf("Expected SectionsPath to return an empty string, but got %s", p.SectionsPath())
	}
}
