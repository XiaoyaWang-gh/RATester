package commands

import (
	"fmt"
	"testing"
)

func TestreplaceImageTag_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple case",
			input:    "{% img class1 src1 100 200 'title1' %}",
			expected: "{{< figure class=\"class1\" src=\"src1\" width=\"100\" height=\"200\" title=\"title1\" >}}",
		},
		{
			name:     "complex case",
			input:    "{% img class2 src2 300 400 'title2' 'alt2' %}",
			expected: "{{< figure class=\"class2\" src=\"src2\" width=\"300\" height=\"400\" title=\"title2\" alt=\"alt2\" >}}",
		},
		{
			name:     "no alt case",
			input:    "{% img class3 src3 500 600 'title3' %}",
			expected: "{{< figure class=\"class3\" src=\"src3\" width=\"500\" height=\"600\" title=\"title3\" >}}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cmd := &importCommand{}
			result := cmd.replaceImageTag(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
