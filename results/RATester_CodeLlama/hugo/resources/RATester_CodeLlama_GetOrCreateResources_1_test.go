package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestGetOrCreateResources_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	key := "key"
	f := func() (resource.Resources, error) {
		return resource.Resources{}, nil
	}
	c := &ResourceCache{}

	// When
	_, err := c.GetOrCreateResources(key, f)

	// Then
	assert.NoError(t, err)
}
