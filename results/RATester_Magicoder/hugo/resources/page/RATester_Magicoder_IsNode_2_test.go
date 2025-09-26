package page

import (
	"fmt"
	"testing"
)

func TestIsNode_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.IsNode() != false {
		t.Errorf("Expected IsNode() to return false, but it didn't")
	}
}
