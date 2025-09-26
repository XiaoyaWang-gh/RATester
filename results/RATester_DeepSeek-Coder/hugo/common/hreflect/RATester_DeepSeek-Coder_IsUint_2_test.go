package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsUint_2(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Uint",
			kind: reflect.Uint,
			want: true,
		},
		{
			name: "Uint8",
			kind: reflect.Uint8,
			want: true,
		},
		{
			name: "Uint16",
			kind: reflect.Uint16,
			want: true,
		},
		{
			name: "Uint32",
			kind: reflect.Uint32,
			want: true,
		},
		{
			name: "Uint64",
			kind: reflect.Uint64,
			want: true,
		},
		{
			name: "Invalid Kind",
			kind: reflect.Invalid,
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

			if got := IsUint(tt.kind); got != tt.want {
				t.Errorf("IsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
