package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/resources/resource_factories/create"
)

func TestByType_3(t *testing.T) {
	ns := &Namespace{
		createClient: &create.Client{},
	}

	testCases := []struct {
		name     string
		typ      any
		expected resource.Resources
	}{
		{
			name:     "valid type",
			typ:      "validType",
			expected: resource.Resources{},
		},
		{
			name:     "invalid type",
			typ:      "invalidType",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := ns.ByType(tc.typ)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
