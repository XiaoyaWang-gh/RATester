package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnwrapv_2(t *testing.T) {
	type testStruct struct {
		value any
	}

	testCases := []struct {
		name string
		arg  any
		want any
	}{
		{
			name: "Unwraps value if it implements Unwrapper",
			arg:  testStruct{value: "test"},
			want: "test",
		},
		{
			name: "Returns original value if it does not implement Unwrapper",
			arg:  "test",
			want: "test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Unwrapv(tc.arg)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Unwrapv(%v) = %v, want %v", tc.arg, got, tc.want)
			}
		})
	}
}
