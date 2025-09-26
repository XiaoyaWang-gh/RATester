package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddFuncs_1(t *testing.T) {
	type testCase struct {
		name string
		out  FuncMap
		in   FuncMap
		want FuncMap
	}

	testCases := []testCase{
		{
			name: "addFuncs_empty",
			out:  FuncMap{},
			in:   FuncMap{},
			want: FuncMap{},
		},
		{
			name: "addFuncs_single",
			out:  FuncMap{"existing": func() {}},
			in:   FuncMap{"new": func() {}},
			want: FuncMap{"existing": func() {}, "new": func() {}},
		},
		{
			name: "addFuncs_multiple",
			out:  FuncMap{"existing1": func() {}, "existing2": func() {}},
			in:   FuncMap{"new1": func() {}, "new2": func() {}},
			want: FuncMap{"existing1": func() {}, "existing2": func() {}, "new1": func() {}, "new2": func() {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			addFuncs(tc.out, tc.in)
			if !reflect.DeepEqual(tc.out, tc.want) {
				t.Errorf("addFuncs(%v, %v) = %v, want %v", tc.out, tc.in, tc.out, tc.want)
			}
		})
	}
}
