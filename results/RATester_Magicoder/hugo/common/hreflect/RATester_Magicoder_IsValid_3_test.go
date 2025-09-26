package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValid_3(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want bool
	}{
		{
			name: "valid value",
			v:    reflect.ValueOf("test"),
			want: true,
		},
		{
			name: "invalid value",
			v:    reflect.Value{},
			want: false,
		},
		{
			name: "nil pointer",
			v:    reflect.ValueOf((*int)(nil)),
			want: false,
		},
		{
			name: "non-nil pointer",
			v:    reflect.ValueOf(&[]int{}),
			want: true,
		},
		{
			name: "nil interface",
			v:    reflect.ValueOf((*interface{})(nil)),
			want: false,
		},
		{
			name: "non-nil interface",
			v:    reflect.ValueOf([]int{}),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsValid(tt.v); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
