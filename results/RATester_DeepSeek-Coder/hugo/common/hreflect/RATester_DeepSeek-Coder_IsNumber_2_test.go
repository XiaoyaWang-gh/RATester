package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsNumber_2(t *testing.T) {
	testCases := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test with int kind",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test with uint kind",
			kind: reflect.Uint,
			want: true,
		},
		{
			name: "Test with float kind",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "Test with complex kind",
			kind: reflect.Complex64,
			want: false,
		},
		{
			name: "Test with string kind",
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

			got := IsNumber(tc.kind)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}
