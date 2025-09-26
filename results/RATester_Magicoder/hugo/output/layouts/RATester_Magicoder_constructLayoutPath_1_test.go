package layouts

import (
	"fmt"
	"testing"
)

func TestconstructLayoutPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tests := []struct {
		typ        string
		layout     string
		variations string
		extension  string
		expected   string
	}{
		{"", "", "", "", ""},
		{"", "layout", "", "", "layout"},
		{"", "", "variations", "", "variations"},
		{"", "", "", "extension", "extension"},
		{"type", "layout", "variations", "extension", "type/layout.variations.extension"},
		{"type", "layout", "", "extension", "type/layout.extension"},
		{"type", "", "variations", "extension", "type/variations.extension"},
		{"type", "", "", "extension", "type/extension"},
		{"", "layout", "variations", "", "layout.variations"},
		{"", "layout", "", "extension", "layout.extension"},
		{"", "", "variations", "extension", "variations.extension"},
		{"", "", "", "", ""},
	}

	for _, test := range tests {
		result := constructLayoutPath(test.typ, test.layout, test.variations, test.extension)
		if result != test.expected {
			t.Errorf("constructLayoutPath(%s, %s, %s, %s) = %s; expected %s", test.typ, test.layout, test.variations, test.extension, result, test.expected)
		}
	}
}
