package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestgetRoot_1(t *testing.T) {
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
			name:     "Existing key",
			key:      "existingKey",
			expected: []RootMapping{{From: "existingKey"}},
		},
		{
			name:     "Non-existing key",
			key:      "nonExistingKey",
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

			fs.rootMapToReal.Insert(tc.key, tc.expected)
			result := fs.getRoot(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
