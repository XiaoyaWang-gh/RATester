package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVarValue_1(t *testing.T) {
	s := &state{
		vars: []variable{
			{name: "var1", value: reflect.ValueOf(1)},
			{name: "var2", value: reflect.ValueOf(2)},
		},
	}

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"var1", "var1", 1},
		{"var2", "var2", 2},
		{"undefined", "undefined", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := s.varValue(test.input)
			if result.Int() != int64(test.expected) {
				t.Errorf("Expected %d, got %d", test.expected, result.Int())
			}
		})
	}
}
