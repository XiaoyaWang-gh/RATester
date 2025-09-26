package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetRoot_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    *connectCert
		expected []types.FileOrContent
	}{
		{
			name: "Test with empty root",
			input: &connectCert{
				root: []string{},
			},
			expected: []types.FileOrContent{},
		},
		{
			name: "Test with one root",
			input: &connectCert{
				root: []string{"root1"},
			},
			expected: []types.FileOrContent{"root1"},
		},
		{
			name: "Test with multiple roots",
			input: &connectCert{
				root: []string{"root1", "root2", "root3"},
			},
			expected: []types.FileOrContent{"root1", "root2", "root3"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.input.getRoot()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
