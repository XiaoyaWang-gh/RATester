package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsFloat_2(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "float32",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "float64",
			kind: reflect.Float64,
			want: true,
		},
		{
			name: "int",
			kind: reflect.Int,
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

			if got := IsFloat(tt.kind); got != tt.want {
				t.Errorf("IsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
