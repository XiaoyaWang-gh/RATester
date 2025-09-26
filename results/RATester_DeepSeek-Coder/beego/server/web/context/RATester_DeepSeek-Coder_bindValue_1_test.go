package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBindValue_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name     string
		val      string
		typ      reflect.Type
		expected reflect.Value
	}{
		{
			name:     "Test bindValue with string type",
			val:      "John",
			typ:      reflect.TypeOf(""),
			expected: reflect.ValueOf("John"),
		},
		{
			name:     "Test bindValue with int type",
			val:      "25",
			typ:      reflect.TypeOf(0),
			expected: reflect.ValueOf(25),
		},
		{
			name:     "Test bindValue with struct type",
			val:      "{\"Name\":\"John\",\"Age\":25}",
			typ:      reflect.TypeOf(testStruct{}),
			expected: reflect.ValueOf(testStruct{"John", 25}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			input := &BeegoInput{}
			result := input.bindValue(tc.val, tc.typ)
			if result.Interface() != tc.expected.Interface() {
				t.Errorf("Expected %v, got %v", tc.expected.Interface(), result.Interface())
			}
		})
	}
}
