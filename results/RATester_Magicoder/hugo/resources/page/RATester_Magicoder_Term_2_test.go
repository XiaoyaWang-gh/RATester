package page

import (
	"fmt"
	"testing"
)

func TestTerm_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ie := OrderedTaxonomyEntry{Name: "Technology"}
	expected := "Technology"
	result := ie.Term()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
