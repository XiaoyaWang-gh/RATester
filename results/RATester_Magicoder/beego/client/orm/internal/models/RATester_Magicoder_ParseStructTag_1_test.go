package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseStructTag_1(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected map[string]bool
	}{
		{
			name: "test case 1",
			data: "field name string",
			expected: map[string]bool{
				"field": true,
			},
		},
		{
			name: "test case 2",
			data: "field name string, field2(tag) string",
			expected: map[string]bool{
				"field":  true,
				"field2": true,
			},
		},
		{
			name: "test case 3",
			data: "field name string, field2(tag) string, field3(tag2) string",
			expected: map[string]bool{
				"field":  true,
				"field2": true,
				"field3": true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			attrs, _ := ParseStructTag(test.data)
			if !reflect.DeepEqual(attrs, test.expected) {
				t.Errorf("Expected: %v, Got: %v", test.expected, attrs)
			}
		})
	}
}
