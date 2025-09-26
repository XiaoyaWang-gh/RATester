package navigation

import (
	"fmt"
	"testing"
)

func TestHasChildren_1(t *testing.T) {
	type testStruct struct {
		name     string
		menu     *MenuEntry
		expected bool
	}

	tests := []testStruct{
		{
			name: "Has Children",
			menu: &MenuEntry{
				Children: make(Menu, 0),
			},
			expected: true,
		},
		{
			name: "No Children",
			menu: &MenuEntry{
				Children: nil,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.menu.HasChildren()
			if result != test.expected {
				t.Errorf("Expected %t, got %t", test.expected, result)
			}
		})
	}
}
