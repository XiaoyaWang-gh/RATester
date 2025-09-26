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
			name: "test case 1",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "test case 2",
			kind: reflect.Uint,
			want: true,
		},
		{
			name: "test case 3",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "test case 4",
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

			if got := isNumber(tt.kind); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
