package postpub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructToMapWithPlaceholders_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name   string
		root   string
		in     any
		expect map[string]any
	}{
		{
			name: "Test with a struct",
			root: "root",
			in: testStruct{
				Name: "John",
				Age:  30,
			},
			expect: map[string]any{
				"root.Name": "John",
				"root.Age":  30,
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

			result := structToMapWithPlaceholders(tc.root, tc.in, func(s string) string {
				return fmt.Sprintf(":%s", s)
			})

			if !reflect.DeepEqual(result, tc.expect) {
				t.Errorf("Expected %v, got %v", tc.expect, result)
			}
		})
	}
}
