package page

import (
	"fmt"
	"testing"
)

func TestUniqueID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.UniqueID() != "" {
		t.Errorf("Expected UniqueID to be empty, but got %s", p.UniqueID())
	}
}
