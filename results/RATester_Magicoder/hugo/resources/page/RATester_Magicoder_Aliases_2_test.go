package page

import (
	"fmt"
	"testing"
)

func TestAliases_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := nopPage(0)
	result := p.Aliases()
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}
}
