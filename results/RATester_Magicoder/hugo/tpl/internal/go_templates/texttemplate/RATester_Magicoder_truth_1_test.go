package template

import (
	"fmt"
	"reflect"
	"testing"
)

func Testtruth_1(t *testing.T) {
	tests := []struct {
		name string
		arg  reflect.Value
		want bool
	}{
		{
			name: "Test case 1",
			arg:  reflect.ValueOf(true),
			want: true,
		},
		{
			name: "Test case 2",
			arg:  reflect.ValueOf(false),
			want: false,
		},
		{
			name: "Test case 3",
			arg:  reflect.ValueOf(0),
			want: false,
		},
		{
			name: "Test case 4",
			arg:  reflect.ValueOf(1),
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

			if got := truth(tt.arg); got != tt.want {
				t.Errorf("truth() = %v, want %v", got, tt.want)
			}
		})
	}
}
