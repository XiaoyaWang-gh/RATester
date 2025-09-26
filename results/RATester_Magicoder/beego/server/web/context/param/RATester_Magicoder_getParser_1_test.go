package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetParser_1(t *testing.T) {
	tests := []struct {
		name   string
		param  *MethodParam
		input  reflect.Type
		output paramParser
	}{
		{
			name: "Test int",
			param: &MethodParam{
				name:     "test",
				in:       body,
				required: true,
			},
			input:  reflect.TypeOf(1),
			output: intParser{},
		},
		{
			name: "Test string",
			param: &MethodParam{
				name:     "test",
				in:       body,
				required: true,
			},
			input:  reflect.TypeOf("test"),
			output: stringParser{},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := getParser(test.param, test.input)
			if reflect.TypeOf(result) != reflect.TypeOf(test.output) {
				t.Errorf("Expected %T, got %T", test.output, result)
			}
		})
	}
}
