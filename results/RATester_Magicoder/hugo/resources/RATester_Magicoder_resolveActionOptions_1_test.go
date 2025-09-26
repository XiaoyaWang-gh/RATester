package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestresolveActionOptions_1(t *testing.T) {
	tests := []struct {
		name     string
		spec     string
		expected string
		options  []string
	}{
		{
			name:     "Test case 1",
			spec:     "crop 100x100",
			expected: "crop",
			options:  []string{"100x100"},
		},
		{
			name:     "Test case 2",
			spec:     "resize 200x200",
			expected: "resize",
			options:  []string{"200x200"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &imageResource{}
			action, options := i.resolveActionOptions(tt.spec)
			if action != tt.expected {
				t.Errorf("Expected action %s, but got %s", tt.expected, action)
			}
			if !reflect.DeepEqual(options, tt.options) {
				t.Errorf("Expected options %v, but got %v", tt.options, options)
			}
		})
	}
}
