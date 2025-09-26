package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestencodeUint_1(t *testing.T) {
	tests := []struct {
		name string
		v    reflect.Value
		want string
	}{
		{
			name: "Test case 1",
			v:    reflect.ValueOf(uint(10)),
			want: "10",
		},
		{
			name: "Test case 2",
			v:    reflect.ValueOf(uint(20)),
			want: "20",
		},
		{
			name: "Test case 3",
			v:    reflect.ValueOf(uint(30)),
			want: "30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := encodeUint(tt.v); got != tt.want {
				t.Errorf("encodeUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
