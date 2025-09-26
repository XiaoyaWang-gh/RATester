package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/resources/resource_factories/create"
)

func TestMatch_3(t *testing.T) {
	testCases := []struct {
		name     string
		pattern  any
		expected resource.Resources
	}{
		{
			name:     "valid pattern",
			pattern:  "*",
			expected: []resource.Resource{
				// add some mock resources here
			},
		},
		{
			name:    "invalid pattern",
			pattern: 123,
			expected: func() resource.Resources {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()

				ns := &Namespace{
					createClient: &create.Client{},
				}

				return ns.Match(123)
			}(),
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

			result := ns.Match(tc.pattern)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
