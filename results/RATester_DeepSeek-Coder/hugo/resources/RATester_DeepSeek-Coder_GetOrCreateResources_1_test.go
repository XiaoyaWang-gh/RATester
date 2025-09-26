package resources

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
	"github.com/zeebo/assert"
)

func TestGetOrCreateResources_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		key      string
		resource resource.Resources
		err      error
	}{
		{
			name:     "success",
			key:      "testKey",
			resource: resource.Resources{},
			err:      nil,
		},
		{
			name:     "failure",
			key:      "failKey",
			resource: nil,
			err:      errors.New("failed to get resources"),
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			cache := &ResourceCache{}

			getOrCreateFunc := func() (resource.Resources, error) {
				return tc.resource, tc.err
			}

			resources, err := cache.GetOrCreateResources(tc.key, getOrCreateFunc)

			if tc.err != nil {
				assert.Error(t, err)
				assert.Nil(t, resources)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.resource, resources)
			}
		})
	}
}
