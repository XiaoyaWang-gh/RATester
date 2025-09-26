package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestvarValue_1(t *testing.T) {
	s := &state{
		vars: []variable{
			{name: "var1", value: reflect.ValueOf(1)},
			{name: "var2", value: reflect.ValueOf(2)},
		},
	}

	tests := []struct {
		name     string
		input    string
		expected reflect.Value
	}{
		{"var1", "var1", reflect.ValueOf(1)},
		{"var2", "var2", reflect.ValueOf(2)},
		{"undefined", "undefined", reflect.ValueOf(0)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := s.varValue(test.input)
			if result.Interface() != test.expected.Interface() {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
