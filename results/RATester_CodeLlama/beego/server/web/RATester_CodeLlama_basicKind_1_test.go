package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasicKind_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want kind
	}{
		{
			name: "bool",
			v:    reflect.ValueOf(true),
			want: boolKind,
		},
		{
			name: "int",
			v:    reflect.ValueOf(1),
			want: intKind,
		},
		{
			name: "uint",
			v:    reflect.ValueOf(uint(1)),
			want: uintKind,
		},
		{
			name: "float",
			v:    reflect.ValueOf(float32(1)),
			want: floatKind,
		},
		{
			name: "string",
			v:    reflect.ValueOf(""),
			want: stringKind,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got, _ := basicKind(tt.v); got != tt.want {
				t.Errorf("basicKind() = %v, want %v", got, tt.want)
			}
		})
	}
}
