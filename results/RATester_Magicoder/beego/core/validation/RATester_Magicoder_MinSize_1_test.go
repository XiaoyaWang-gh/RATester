package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Validation{}

	type testCase struct {
		obj  interface{}
		min  int
		key  string
		want *Result
	}

	testCases := []testCase{
		{"", 1, "key", &Result{Error: &Error{}, Ok: false}},
		{"a", 1, "key", &Result{Error: &Error{}, Ok: true}},
		{"ab", 2, "key", &Result{Error: &Error{}, Ok: true}},
		{"abc", 3, "key", &Result{Error: &Error{}, Ok: true}},
		{"abcd", 4, "key", &Result{Error: &Error{}, Ok: true}},
		{"abcde", 5, "key", &Result{Error: &Error{}, Ok: true}},
		{"abcdef", 6, "key", &Result{Error: &Error{}, Ok: true}},
	}

	for _, tc := range testCases {
		got := v.MinSize(tc.obj, tc.min, tc.key)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("MinSize(%v, %v, %v) = %v, want %v", tc.obj, tc.min, tc.key, got, tc.want)
		}
	}
}
