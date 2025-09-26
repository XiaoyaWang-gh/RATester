package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFieldsSlash_1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  []string
	}{
		{"empty string", "", []string{}},
		{"single field", "field", []string{"field"}},
		{"multiple fields", "field1/field2/field3", []string{"field1", "field2", "field3"}},
		{"leading slash", "/field1/field2/field3", []string{"field1", "field2", "field3"}},
		{"trailing slash", "field1/field2/field3/", []string{"field1", "field2", "field3"}},
		{"multiple slashes", "field1//field2/field3", []string{"field1", "field2", "field3"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := FieldsSlash(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
