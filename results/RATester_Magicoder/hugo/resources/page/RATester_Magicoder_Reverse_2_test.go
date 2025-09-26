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

	reversedTaxonomy := taxonomy.Reverse()

	if reversedTaxonomy[0].Name != "Entry3" {
		t.Errorf("Expected first entry to be 'Entry3', but got %s", reversedTaxonomy[0].Name)
	}

	if reversedTaxonomy[2].Name != "Entry1" {
		t.Errorf("Expected last entry to be 'Entry1', but got %s", reversedTaxonomy[2].Name)
	}
}
