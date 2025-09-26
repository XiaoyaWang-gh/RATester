package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeInt_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(1),
			want: "1",
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(2),
			want: "2",
		},
		{
			name: "Test case 3",
			v:    reflect.ValueOf(3),
			want: "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeInt(tt.v); got != tt.want {
				t.Errorf("encodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
