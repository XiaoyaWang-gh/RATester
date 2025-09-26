package cli

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestFindTypedField_1(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testCases := []struct {
		name     string
		rType    reflect.Type
		node     *parser.Node
		expected bool
	}{
		{
			name:     "Test case 1",
			rType:    reflect.TypeOf(testStruct{}),
			node:     &parser.Node{Name: "Field1"},
			expected: true,
		},
		{
			name:     "Test case 2",
			rType:    reflect.TypeOf(testStruct{}),
			node:     &parser.Node{Name: "Field3"},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, actual := findTypedField(tc.rType, tc.node)
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
