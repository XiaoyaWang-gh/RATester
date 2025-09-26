package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMethodIndexByName_1(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 int
	}

	testCases := []struct {
		name   string
		tp     reflect.Type
		method string
		want   int
	}{
		{
			name:   "Test method exists",
			tp:     reflect.TypeOf(TestStruct{}),
			method: "Field1",
			want:   0,
		},
		{
			name:   "Test method does not exist",
			tp:     reflect.TypeOf(TestStruct{}),
			method: "NonExistentMethod",
			want:   -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := GetMethodIndexByName(tc.tp, tc.method)
			if got != tc.want {
				t.Errorf("GetMethodIndexByName(%v, %q) = %v, want %v", tc.tp, tc.method, got, tc.want)
			}
		})
	}
}
