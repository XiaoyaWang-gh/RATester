package param

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestParse_11(t *testing.T) {
	t.Parallel()

	parser := jsonParser{}

	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name     string
		value    string
		toType   reflect.Type
		expected interface{}
		err      error
	}{
		{
			name:     "success",
			value:    `{"field":"test"}`,
			toType:   reflect.TypeOf(testStruct{}),
			expected: testStruct{Field: "test"},
			err:      nil,
		},
		{
			name:     "failure",
			value:    `{"field":123}`,
			toType:   reflect.TypeOf(testStruct{}),
			expected: nil,
			err:      errors.New("json: cannot unmarshal number into Go struct field testStruct.Field of type string"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := parser.parse(tc.value, tc.toType)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
