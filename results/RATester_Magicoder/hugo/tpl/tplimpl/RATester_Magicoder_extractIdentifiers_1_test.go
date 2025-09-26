package tplimpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestextractIdentifiers_1(t *testing.T) {
	handler := &templateHandler{}

	tests := []struct {
		name     string
		line     string
		expected []string
	}{
		{
			name:     "simple case",
			line:     "{{< partial \"test\" >}}",
			expected: []string{"test"},
		},
		{
			name:     "multiple identifiers",
			line:     "{{< partial \"test\" >}} {{< partial \"another\" >}}",
			expected: []string{"test", "another"},
		},
		{
			name:     "no identifiers",
			line:     "{{< partial \"\" >}}",
			expected: []string{""},
		},
		{
			name:     "no partial",
			line:     "{{< no_partial \"test\" >}}",
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			identifiers := handler.extractIdentifiers(test.line)
			if !reflect.DeepEqual(identifiers, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, identifiers)
			}
		})
	}
}
