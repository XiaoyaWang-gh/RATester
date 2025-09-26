package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/resource_factories/create"
)

func TestGetMatch_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	pattern := "foo"
	ns := &Namespace{
		createClient: &create.Client{},
	}

	// When
	r := ns.GetMatch(pattern)

	// Then
	assert.NotNil(t, r)
}
