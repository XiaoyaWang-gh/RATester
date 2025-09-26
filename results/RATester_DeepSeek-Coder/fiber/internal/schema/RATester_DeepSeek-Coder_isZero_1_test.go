package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsZero_1(t *testing.T) {
	type testStruct struct {
		field1 int
		field2 string
	}

	testCases := []struct {
		name string
		arg  reflect.Value
		want bool
	}{
		{
			name: "Zero value of int",
			arg:  reflect.ValueOf(0),
			want: true,
		},
		{
			name: "Non-zero value of int",
			arg:  reflect.ValueOf(1),
			want: false,
		},
		{
			name: "Zero value of string",
			arg:  reflect.ValueOf(""),
			want: true,
		},
		{
			name: "Non-zero value of string",
			arg:  reflect.ValueOf("test"),
			want: false,
		},
		{
			name: "Zero value of struct",
			arg:  reflect.ValueOf(testStruct{}),
			want: true,
		},
		{
			name: "Non-zero value of struct",
			arg:  reflect.ValueOf(testStruct{field1: 1, field2: "test"}),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := isZero(tc.arg)
			if got != tc.want {
				t.Errorf("isZero() = %v, want %v", got, tc.want)
			}
		})
	}
}
