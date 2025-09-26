package page

import (
	"fmt"
	"testing"
)

func TestTaxonomies_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	site := testSite{}
	taxonomies := site.Taxonomies()
	if taxonomies == nil {
		t.Error("Expected non-nil taxonomies, but got nil")
	}
}
