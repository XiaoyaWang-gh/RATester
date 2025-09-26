package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTaxonomyArray_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := Taxonomy{"a": WeightedPages{}, "b": WeightedPages{}}
	expected := OrderedTaxonomy{
		OrderedTaxonomyEntry{Name: "a", WeightedPages: WeightedPages{}},
		OrderedTaxonomyEntry{Name: "b", WeightedPages: WeightedPages{}},
	}
	if !reflect.DeepEqual(i.TaxonomyArray(), expected) {
		t.Errorf("TaxonomyArray() = %v, want %v", i.TaxonomyArray(), expected)
	}
}
