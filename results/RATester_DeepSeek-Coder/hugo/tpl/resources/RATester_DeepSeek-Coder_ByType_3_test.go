package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/resources/resource_factories/create"
)

func TestByType_3(t *testing.T) {
	testCases := []struct {
		name     string
		typ      any
		expected resource.Resources
	}{
		{
			name:     "Test with string type",
			typ:      "string",
			expected: resource.Resources{},
		},
		{
			name:     "Test with int type",
			typ:      123,
			expected: resource.Resources{},
		},
		{
			name:     "Test with float type",
			typ:      123.456,
			expected: resource.Resources{},
		},
		{
			name:     "Test with bool type",
			typ:      true,
			expected: resource.Resources{},
		},
		{
			name:     "Test with nil type",
			typ:      nil,
			expected: resource.Resources{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{
				createClient: &create.Client{},
			}

			result := ns.ByType(tc.typ)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
