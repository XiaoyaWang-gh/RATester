package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestisFloat_1(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Float32",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "Float64",
			kind: reflect.Float64,
			want: true,
		},
		{
			name: "Int",
			kind: reflect.Int,
			want: false,
		},
		{
			name: "String",
			kind: reflect.String,
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

			if got := isFloat(tt.kind); got != tt.want {
				t.Errorf("isFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
