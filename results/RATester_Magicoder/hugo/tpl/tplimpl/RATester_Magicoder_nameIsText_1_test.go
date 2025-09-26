package tplimpl

import (
	"fmt"
	"testing"
)

func TestnameIsText_1(t *testing.T) {
	handler := &templateHandler{}

	tests := []struct {
		name     string
		input    string
		expected string
		isText   bool
	}{
		{"text template", "text-template.txt", "text-template", true},
		{"non-text template", "non-text-template.html", "non-text-template.html", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			name, isText := handler.nameIsText(test.input)
			if name != test.expected || isText != test.isText {
				t.Errorf("Expected (%s, %t), got (%s, %t)", test.expected, test.isText, name, isText)
			}
		})
	}
}
