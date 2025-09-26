package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestencodeBool_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(true),
			want: "true",
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(false),
			want: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeBool(tt.v); got != tt.want {
				t.Errorf("encodeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
