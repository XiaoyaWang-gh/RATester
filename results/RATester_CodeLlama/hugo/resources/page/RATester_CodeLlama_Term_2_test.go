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
	if ie.Term() != "Technology" {
		t.Errorf("Term() = %v, want %v", ie.Term(), "Technology")
	}
}
