package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetVar_1(t *testing.T) {
	s := &state{
		vars: []variable{
			{name: "var1", value: reflect.ValueOf(1)},
			{name: "var2", value: reflect.ValueOf(2)},
		},
	}

	testCases := []struct {
		name     string
		input    string
		value    reflect.Value
		expected []variable
	}{
		{
			name:  "existing variable",
			input: "var1",
			value: reflect.ValueOf(3),
			expected: []variable{
				{name: "var1", value: reflect.ValueOf(3)},
				{name: "var2", value: reflect.ValueOf(2)},
			},
		},
		{
			name:  "new variable",
			input: "var3",
			value: reflect.ValueOf(4),
			expected: []variable{
				{name: "var1", value: reflect.ValueOf(1)},
				{name: "var2", value: reflect.ValueOf(2)},
				{name: "var3", value: reflect.ValueOf(4)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s.setVar(tc.input, tc.value)
			if !reflect.DeepEqual(s.vars, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, s.vars)
			}
		})
	}
}
