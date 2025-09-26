package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsEmptyFields_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testCase struct {
		fields []fieldWithPrefix
		src    map[string][]string
		want   bool
	}

	testCases := []testCase{
		{
			fields: []fieldWithPrefix{
				{
					fieldInfo: &fieldInfo{
						typ: reflect.TypeOf(""),
					},
					prefix: "",
				},
			},
			src: map[string][]string{
				"field1": {"value1"},
			},
			want: false,
		},
		{
			fields: []fieldWithPrefix{
				{
					fieldInfo: &fieldInfo{
						typ: reflect.TypeOf(""),
					},
					prefix: "",
				},
			},
			src: map[string][]string{
				"field1": {},
			},
			want: true,
		},
		{
			fields: []fieldWithPrefix{
				{
					fieldInfo: &fieldInfo{
						typ: reflect.TypeOf(""),
					},
					prefix: "",
				},
			},
			src: map[string][]string{
				"field1": {"value1", "value2"},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		got := isEmptyFields(tc.fields, tc.src)
		if got != tc.want {
			t.Errorf("isEmptyFields(%v, %v) = %v, want %v", tc.fields, tc.src, got, tc.want)
		}
	}
}
