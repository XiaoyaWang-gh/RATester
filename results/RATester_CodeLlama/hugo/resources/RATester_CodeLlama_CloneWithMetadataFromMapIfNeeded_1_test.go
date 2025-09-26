package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestCloneWithMetadataFromMapIfNeeded_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var m []map[string]any
	var r resource.Resource

	// when
	result := CloneWithMetadataFromMapIfNeeded(m, r)

	// then
	assert.Equal(t, r, result)
}
