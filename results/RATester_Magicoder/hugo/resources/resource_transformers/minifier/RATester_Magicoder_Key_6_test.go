package minifier

import (
	"fmt"
	"testing"
)

func TestKey_6(t *testing.T) {
	tt := []struct {
		name     string
		expected string
	}{
		{
			name:     "minify",
			expected: "minify",
		},
		{
			name:     "gzip",
			expected: "gzip",
		},
		{
			name:     "uglify",
			expected: "uglify",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			minify := &minifyTransformation{}
			key := minify.Key()
			if key.Value() != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, key.Value())
			}
		})
	}
}
