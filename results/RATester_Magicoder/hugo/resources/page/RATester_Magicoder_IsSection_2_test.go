package page

import (
	"fmt"
	"testing"
)

func TestIsSection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsSection() != false {
		t.Errorf("Expected IsSection to return false, but it did not.")
	}
}
