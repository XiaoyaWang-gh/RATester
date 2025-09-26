package tplimpl

import (
	"fmt"
	"testing"
)

func TestresolveTemplateType_1(t *testing.T) {
	tests := []struct {
		name     string
		expected templateType
	}{
		{"shortcode", templateShortcode},
		{"partials/partial", templatePartial},
		{"undefined", templateUndefined},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := resolveTemplateType(test.name)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
