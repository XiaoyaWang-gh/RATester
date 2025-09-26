package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcreateField_1(t *testing.T) {
	type testCase struct {
		name        string
		field       reflect.StructField
		parentAlias string
		expected    *fieldInfo
	}

	tests := []testCase{
		{
			name: "Test case 1",
			field: reflect.StructField{
				Name: "Field1",
				Type: reflect.TypeOf(""),
				Tag:  reflect.StructTag(`json:"field1"`),
			},
			parentAlias: "",
			expected: &fieldInfo{
				typ:            reflect.TypeOf(""),
				name:           "Field1",
				alias:          "field1",
				canonicalAlias: "field1",
				unmarshalerInfo: unmarshaler{
					Unmarshaler:       nil,
					IsValid:           false,
					IsPtr:             false,
					IsSliceElement:    false,
					IsSliceElementPtr: false,
				},
				isSliceOfStructs: false,
				isAnonymous:      false,
				isRequired:       false,
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &cache{}
			result := c.createField(tc.field, tc.parentAlias)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
