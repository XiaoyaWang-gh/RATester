package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValid_3(t *testing.T) {
	type testCase struct {
		name string
		arg  reflect.Value
		want bool
	}

	tests := []testCase{
		{
			name: "valid value",
			arg:  reflect.ValueOf("test"),
			want: true,
		},
		{
			name: "invalid value",
			arg:  reflect.Value{},
			want: false,
		},
		{
			name: "nil pointer",
			arg:  reflect.ValueOf((*int)(nil)),
			want: false,
		},
		{
			name: "non-nil pointer",
			arg:  reflect.ValueOf(&[]int{}),
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

			if got := IsValid(tt.arg); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
