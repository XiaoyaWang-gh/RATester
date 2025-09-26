package tplimpl

import (
	"fmt"
	"testing"
)

func TestgetIfNotVisited_1(t *testing.T) {
	ctx := &templateContext{
		visited:          make(map[string]bool),
		templateNotFound: make(map[string]bool),
		lookupFn: func(name string) *templateState {
			if name == "existingTemplate" {
				return &templateState{}
			}
			return nil
		},
	}

	tests := []struct {
		name     string
		ctx      *templateContext
		expected *templateState
	}{
		{
			name:     "existingTemplate",
			ctx:      ctx,
			expected: &templateState{},
		},
		{
			name: "nonExistingTemplate",
			ctx:  ctx,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.ctx.getIfNotVisited(test.name)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
