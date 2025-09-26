package page

import (
	"fmt"
	"testing"
)

func TestWeight_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	if p.Weight() != 0 {
		t.Errorf("Expected Weight() to return 0, but got %d", p.Weight())
	}
}
