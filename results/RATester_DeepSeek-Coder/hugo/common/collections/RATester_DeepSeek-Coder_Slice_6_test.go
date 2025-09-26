package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSlice_6(t *testing.T) {
	type testCase struct {
		name string
		args []any
		want any
	}

	testCases := []testCase{
		{
			name: "Test with no arguments",
			args: []any{},
			want: []any{},
		},
		{
			name: "Test with one argument",
			args: []any{1},
			want: []any{1},
		},
		{
			name: "Test with multiple arguments of the same type",
			args: []any{1, 2, 3},
			want: []any{1, 2, 3},
		},
		{
			name: "Test with multiple arguments of different types",
			args: []any{1, "2"},
			want: []any{1, "2"},
		},
		{
			name: "Test with a Slicer",
			args: []any{[]int{1, 2, 3}},
			want: []any{[]int{1, 2, 3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Slice(tc.args...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Slice() = %v, want %v", got, tc.want)
			}
		})
	}
}
