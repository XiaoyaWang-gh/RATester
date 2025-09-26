package context

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParseFormToStruct_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testCase struct {
		form url.Values
		obj  interface{}
		want interface{}
	}
	testCases := []testCase{
		{
			form: url.Values{
				"Name": []string{"test"},
			},
			obj: &struct {
				Name string `form:"Name"`
			}{},
			want: &struct {
				Name string `form:"Name"`
			}{Name: "test"},
		},
	}
	for _, tc := range testCases {
		err := parseFormToStruct(tc.form, reflect.TypeOf(tc.obj), reflect.ValueOf(tc.obj))
		if err != nil {
			t.Errorf("parseFormToStruct error: %v", err)
		}
		if !reflect.DeepEqual(tc.obj, tc.want) {
			t.Errorf("parseFormToStruct error: %v, want: %v", tc.obj, tc.want)
		}
	}
}
