package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestisInt_1(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test case 1: Kind is Int",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test case 2: Kind is Int8",
			kind: reflect.Int8,
			want: true,
		},
		{
			name: "Test case 3: Kind is Int16",
			kind: reflect.Int16,
			want: true,
		},
		{
			name: "Test case 4: Kind is Int32",
			kind: reflect.Int32,
			want: true,
		},
		{
			name: "Test case 5: Kind is Int64",
			kind: reflect.Int64,
			want: true,
		},
		{
			name: "Test case 6: Kind is not Int",
			kind: reflect.Float64,
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

			if got := isInt(tt.kind); got != tt.want {
				t.Errorf("isInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
