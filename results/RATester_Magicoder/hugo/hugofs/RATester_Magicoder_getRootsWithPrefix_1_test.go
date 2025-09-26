package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestgetRootsWithPrefix_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
	}

	testCases := []struct {
		name     string
		prefix   string
		expected []RootMapping
	}{
		{
			name:     "empty prefix",
			prefix:   "",
			expected: []RootMapping{},
		},
		{
			name:     "non-existent prefix",
			prefix:   "nonexistent",
			expected: []RootMapping{},
		},
		{
			name:   "existing prefix",
			prefix: "existing",
			expected: []RootMapping{
				{From: "existing1", To: "real1"},
				{From: "existing2", To: "real2"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			for _, rm := range tc.expected {
				fs.rootMapToReal.Insert(rm.From, rm)
			}

			result := fs.getRootsWithPrefix(tc.prefix)

			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d roots, got %d", len(tc.expected), len(result))
			}

			for i, rm := range result {
				if rm.From != tc.expected[i].From || rm.To != tc.expected[i].To {
					t.Errorf("Expected %v, got %v", tc.expected[i], rm)
				}
			}
		})
	}
}
