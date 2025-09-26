package page

import (
	"fmt"
	"testing"
)

func TestReverse_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	taxonomy := OrderedTaxonomy{
		{Name: "Entry1"},
		{Name: "Entry2"},
		{Name: "Entry3"},
	}

	reversed := taxonomy.Reverse()

	if reversed[0].Name != "Entry3" {
		t.Errorf("Expected first entry to be 'Entry3', got %s", reversed[0].Name)
	}

	if reversed[1].Name != "Entry2" {
		t.Errorf("Expected second entry to be 'Entry2', got %s", reversed[1].Name)
	}

	if reversed[2].Name != "Entry1" {
		t.Errorf("Expected third entry to be 'Entry1', got %s", reversed[2].Name)
	}
}
