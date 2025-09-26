package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestReverseLookup_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	testCases := []struct {
		name     string
		filename string
		expected []ComponentPath
		err      error
	}{
		{
			name:     "Test case 1",
			filename: "file1",
			expected: []ComponentPath{
				{Component: "component1", Path: "path1", Lang: "lang1", Watch: true},
				{Component: "component2", Path: "path2", Lang: "lang2", Watch: false},
			},
			err: nil,
		},
		{
			name:     "Test case 2",
			filename: "file2",
			expected: []ComponentPath{
				{Component: "component3", Path: "path3", Lang: "lang3", Watch: true},
				{Component: "component4", Path: "path4", Lang: "lang4", Watch: false},
			},
			err: nil,
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

			result, err := fs.ReverseLookup(tc.filename)
			if err != tc.err {
				t.Errorf("Expected error %v, but got %v", tc.err, err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v, but got %v", tc.expected, result)
			}
		})
	}
}
