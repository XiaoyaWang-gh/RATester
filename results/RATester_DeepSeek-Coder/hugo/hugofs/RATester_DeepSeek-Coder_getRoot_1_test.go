package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestGetRoot_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
	}

	testCases := []struct {
		name     string
		key      string
		expected []RootMapping
	}{
		{
			name:     "Key exists",
			key:      "testKey",
			expected: []RootMapping{{From: "testFrom"}},
		},
		{
			name:     "Key does not exist",
			key:      "nonExistentKey",
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

			if tc.expected != nil {
				fs.rootMapToReal.Insert(tc.key, tc.expected)
			}

			result := fs.getRoot(tc.key)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
