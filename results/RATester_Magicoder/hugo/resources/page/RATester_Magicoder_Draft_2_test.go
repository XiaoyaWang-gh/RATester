package page

import (
	"fmt"
	"testing"
)

func TestDraft_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Draft() != false {
		t.Errorf("Expected Draft() to return false, but it didn't")
	}
}
