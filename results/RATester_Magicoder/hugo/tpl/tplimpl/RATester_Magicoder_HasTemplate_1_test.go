package tplimpl

import (
	"fmt"
	"testing"
)

func TestHasTemplate_1(t *testing.T) {
	handler := &templateHandler{
		baseof: map[string]templateInfo{
			"base": {name: "base", template: "base template"},
		},
		needsBaseof: map[string]templateInfo{
			"needsBase": {name: "needsBase", template: "needs base template"},
		},
	}

	tests := []struct {
		name     string
		template string
		expected bool
	}{
		{"base", "base", true},
		{"needsBase", "needsBase", true},
		{"notFound", "notFound", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := handler.HasTemplate(test.name)
			if result != test.expected {
				t.Errorf("Expected %t, but got %t", test.expected, result)
			}
		})
	}
}
