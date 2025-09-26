package page

import (
	"fmt"
	"testing"
)

func TestIsHtmlIndex_1(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		expected bool
	}{
		{
			name:     "Test case 1",
			elements: []string{"index.html"},
			expected: true,
		},
		{
			name:     "Test case 2",
			elements: []string{"index.htm"},
			expected: false,
		},
		{
			name:     "Test case 3",
			elements: []string{"index"},
			expected: false,
		},
		{
			name:     "Test case 4",
			elements: []string{"index.html", "subpage.html"},
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

			p := &pagePathBuilder{els: test.elements}
			result := p.IsHtmlIndex()
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
