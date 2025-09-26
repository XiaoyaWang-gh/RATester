package page

import (
	"fmt"
	"testing"
)

func TestIsAncestor_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsAncestor(nil) != false {
		t.Errorf("Expected IsAncestor to return false, but it did not")
	}
}
