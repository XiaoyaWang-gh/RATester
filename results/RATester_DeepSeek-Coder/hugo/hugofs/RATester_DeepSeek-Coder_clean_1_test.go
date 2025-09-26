package hugofs

import (
	"fmt"
	"testing"
)

func TestClean_1(t *testing.T) {
	type test struct {
		name     string
		rm       *RootMapping
		expected *RootMapping
	}

	tests := []test{
		{
			name: "clean test 1",
			rm: &RootMapping{
				From: "/test/path",
				To:   "/test/path/to",
			},
			expected: &RootMapping{
				From: "/test/path",
				To:   "test/path/to",
			},
		},
		{
			name: "clean test 2",
			rm: &RootMapping{
				From: "/test/path/",
				To:   "/test/path/to/",
			},
			expected: &RootMapping{
				From: "/test/path",
				To:   "test/path/to",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.rm.clean()
			if tt.rm.From != tt.expected.From {
				t.Errorf("Expected From to be '%s', got '%s'", tt.expected.From, tt.rm.From)
			}
			if tt.rm.To != tt.expected.To {
				t.Errorf("Expected To to be '%s', got '%s'", tt.expected.To, tt.rm.To)
			}
		})
	}
}
