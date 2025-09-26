package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsNumber_1(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test int kind",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test uint kind",
			kind: reflect.Uint,
			want: true,
		},
		{
			name: "Test float kind",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "Test complex kind",
			kind: reflect.Complex64,
			want: false,
		},
		{
			name: "Test string kind",
			kind: reflect.String,
			want: false,
		},
		{
			name: "Test bool kind",
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

			if got := isNumber(tt.kind); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
