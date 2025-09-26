package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsNumber_2(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test case 1: Kind is int",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test case 2: Kind is uint",
			kind: reflect.Uint,
			want: true,
		},
		{
			name: "Test case 3: Kind is float32",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "Test case 4: Kind is float64",
			kind: reflect.Float64,
			want: true,
		},
		{
			name: "Test case 5: Kind is complex64",
			kind: reflect.Complex64,
			want: false,
		},
		{
			name: "Test case 6: Kind is complex128",
			kind: reflect.Complex128,
			want: false,
		},
		{
			name: "Test case 7: Kind is string",
			kind: reflect.String,
			want: false,
		},
		{
			name: "Test case 8: Kind is bool",
			kind: reflect.Bool,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsNumber(tt.kind); got != tt.want {
				t.Errorf("IsNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
