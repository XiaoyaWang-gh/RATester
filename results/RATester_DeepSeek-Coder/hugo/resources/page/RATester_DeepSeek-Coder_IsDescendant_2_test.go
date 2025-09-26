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

	var p nopPage
	result := p.IsDescendant(nil)
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}
}
