package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/resources/page"
)

func TestGetOrCreatePagesFromCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var m pageMap
	var cache *dynacache.Partition[string, page.Pages]
	var key string
	var create func(string) (page.Pages, error)
	var expected page.Pages
	var actual page.Pages
	var err error

	// Act
	actual, err = m.getOrCreatePagesFromCache(cache, key, create)

	// Assert
	if err != nil {
		t.Errorf("Error was not expected while getting or creating pages from cache: %s", err)
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected pages %v, but got %v", expected, actual)
	}
}
