package gin

import (
	"fmt"
	"testing"
)

func TestCountParams_1(t *testing.T) {
	type test struct {
		name     string
		path     string
		expected uint16
	}

	tests := []test{
		{
			name:     "Test case 1",
			path:     "/user/:id",
			expected: 1,
		},
		{
			name:     "Test case 2",
			path:     "/user/*",
			expected: 1,
		},
		{
			name:     "Test case 3",
			path:     "/user/:id/post/*",
			expected: 2,
		},
		{
			name:     "Test case 4",
			path:     "/user/post/:id",
			expected: 1,
		},
		{
			name:     "Test case 5",
			path:     "/user/post/:id/comment/:id",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := countParams(tt.path)
			if got != tt.expected {
				t.Errorf("countParams(%s) = %d; want %d", tt.path, got, tt.expected)
			}
		})
	}
}
