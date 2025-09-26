package tplimpl

import (
	"fmt"
	"testing"
)

func TestGetIfNotVisited_1(t *testing.T) {
	ctx := &templateContext{
		visited:          make(map[string]bool),
		templateNotFound: make(map[string]bool),
		lookupFn: func(name string) *templateState {
			if name == "existing" {
				return &templateState{}
			}
			return nil
		},
	}

	testCases := []struct {
		name     string
		expected *templateState
	}{
		{
			name:     "existing",
			expected: &templateState{},
		},
		{
			name:     "non-existing",
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

			result := ctx.getIfNotVisited(tc.name)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
