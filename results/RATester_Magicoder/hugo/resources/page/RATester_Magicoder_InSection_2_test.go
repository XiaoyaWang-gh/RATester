package page

import (
	"fmt"
	"testing"
)

func TestInSection_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.InSection(nil) != false {
		t.Errorf("Expected InSection to return false, but it didn't")
	}
}
