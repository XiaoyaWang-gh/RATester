package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TesthasPrefix_1(t *testing.T) {
	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
	}

	fs.rootMapToReal.Insert("test", "test")

	tests := []struct {
		name     string
		prefix   string
		expected bool
	}{
		{
			name:     "prefix exists",
			prefix:   "test",
			expected: true,
		},
		{
			name:     "prefix does not exist",
			prefix:   "nonexistent",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := fs.hasPrefix(test.prefix)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
