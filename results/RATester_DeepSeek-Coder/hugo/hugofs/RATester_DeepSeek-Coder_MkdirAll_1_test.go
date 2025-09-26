package hugofs

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestMkdirAll_1(t *testing.T) {
	fs := noOpFs{}

	testCases := []struct {
		name     string
		path     string
		perm     os.FileMode
		expected error
	}{
		{
			name:     "success",
			path:     "test",
			perm:     0755,
			expected: nil,
		},
		{
			name:     "failure",
			path:     "",
			perm:     0755,
			expected: errors.New("invalid path"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := fs.MkdirAll(tc.path, tc.perm)
			if err != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, err)
			}
		})
	}
}
