package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestReverseLookupComponent_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	testCases := []struct {
		name      string
		component string
		filename  string
		expected  []ComponentPath
	}{
		{
			name:      "Test case 1",
			component: "component1",
			filename:  "file1",
			expected: []ComponentPath{
				{
					Component: "component1",
					Path:      "/path/to/file1",
					Lang:      "lang1",
					Watch:     true,
				},
			},
		},
		{
			name:      "Test case 2",
			component: "component2",
			filename:  "file2",
			expected: []ComponentPath{
				{
					Component: "component2",
					Path:      "/path/to/file2",
					Lang:      "lang2",
					Watch:     false,
				},
			},
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

			result, err := fs.ReverseLookupComponent(tc.component, tc.filename)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
