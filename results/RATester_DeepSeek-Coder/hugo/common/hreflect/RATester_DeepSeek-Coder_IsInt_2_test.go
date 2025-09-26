package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsInt_2(t *testing.T) {
	testCases := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test with int",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test with int8",
			kind: reflect.Int8,
			want: true,
		},
		{
			name: "Test with int16",
			kind: reflect.Int16,
			want: true,
		},
		{
			name: "Test with int32",
			kind: reflect.Int32,
			want: true,
		},
		{
			name: "Test with int64",
			kind: reflect.Int64,
			want: true,
		},
		{
			name: "Test with other kind",
			kind: reflect.String,
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

			got := IsInt(tc.kind)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
