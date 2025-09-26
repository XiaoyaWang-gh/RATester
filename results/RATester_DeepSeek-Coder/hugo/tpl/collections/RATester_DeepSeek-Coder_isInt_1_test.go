package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsInt_1(t *testing.T) {
	tests := []struct {
		name string
		kind reflect.Kind
		want bool
	}{
		{
			name: "Test case 1",
			kind: reflect.Int,
			want: true,
		},
		{
			name: "Test case 2",
			kind: reflect.Int8,
			want: true,
		},
		{
			name: "Test case 3",
			kind: reflect.Int16,
			want: true,
		},
		{
			name: "Test case 4",
			kind: reflect.Int32,
			want: true,
		},
		{
			name: "Test case 5",
			kind: reflect.Int64,
			want: true,
		},
		{
			name: "Test case 6",
			kind: reflect.Float32,
			want: false,
		},
		{
			name: "Test case 7",
			kind: reflect.Float64,
			want: false,
		},
		{
			name: "Test case 8",
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

			if got := isInt(tt.kind); got != tt.want {
				t.Errorf("isInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
