package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxSize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := Validation{}

	type testCase struct {
		obj  interface{}
		max  int
		key  string
		want *Result
	}

	testCases := []testCase{
		{"abc", 3, "key", &Result{nil, true}},
		{"abc", 2, "key", &Result{nil, false}},
		{"abc", 1, "key", &Result{nil, false}},
		{"abc", 0, "key", &Result{nil, false}},
		{"", 0, "key", &Result{nil, true}},
		{"", 1, "key", &Result{nil, true}},
		{"", 2, "key", &Result{nil, false}},
		{"", 3, "key", &Result{nil, false}},
		{nil, 0, "key", &Result{nil, true}},
		{nil, 1, "key", &Result{nil, true}},
		{nil, 2, "key", &Result{nil, false}},
		{nil, 3, "key", &Result{nil, false}},
	}

	for _, tc := range testCases {
		got := v.MaxSize(tc.obj, tc.max, tc.key)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("MaxSize(%v, %v, %v) = %v, want %v", tc.obj, tc.max, tc.key, got, tc.want)
		}
	}
}
