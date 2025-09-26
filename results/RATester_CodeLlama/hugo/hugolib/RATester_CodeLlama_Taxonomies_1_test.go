package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/page"
)

func TestTaxonomies_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	s.init.taxonomies.Do(context.Background())
	assert.Equal(t, page.TaxonomyList{}, s.taxonomies)
}
