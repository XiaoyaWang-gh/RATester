package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestisNumber_1(t *testing.T) {
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
			name: "Test case 3: Kind is float",
			kind: reflect.Float32,
			want: true,
		},
		{
			name: "Test case 4: Kind is not number",
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
