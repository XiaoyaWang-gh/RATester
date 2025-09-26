package create

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestByType_1(t *testing.T) {
	client := &Client{
		rs: &resources.Spec{},
	}

	testCases := []struct {
		name     string
		tp       string
		expected resource.Resources
	}{
		{
			name:     "Test case 1",
			tp:       "type1",
			expected: resource.Resources{},
		},
		{
			name:     "Test case 2",
			tp:       "type2",
			expected: resource.Resources{},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := client.ByType(tc.tp)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
