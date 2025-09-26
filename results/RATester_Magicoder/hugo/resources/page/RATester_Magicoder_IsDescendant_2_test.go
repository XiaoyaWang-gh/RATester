package page

import (
	"fmt"
	"testing"
)

func TestIsDescendant_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsDescendant(nil) != false {
		t.Errorf("Expected IsDescendant to return false, but got true")
	}
}
