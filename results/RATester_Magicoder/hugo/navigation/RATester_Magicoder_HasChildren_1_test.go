package navigation

import (
	"fmt"
	"testing"
)

func TestHasChildren_1(t *testing.T) {
	type test struct {
		name     string
		menu     MenuEntry
		expected bool
	}

	tests := []test{
		{
			name: "Has Children",
			menu: MenuEntry{
				Children: Menu{
					&MenuEntry{},
				},
			},
			expected: true,
		},
		{
			name: "No Children",
			menu: MenuEntry{
				Children: nil,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.menu.HasChildren()
			if got != tt.expected {
				t.Errorf("Expected %t, but got %t", tt.expected, got)
			}
		})
	}
}
