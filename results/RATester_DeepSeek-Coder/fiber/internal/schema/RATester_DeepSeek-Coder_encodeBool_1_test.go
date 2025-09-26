package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeBool_1(t *testing.T) {
	type testCase struct {
		name string
		arg  reflect.Value
		want string
	}

	tests := []testCase{
		{
			name: "Test case 1",
			arg:  reflect.ValueOf(true),
			want: "true",
		},
		{
			name: "Test case 2",
			arg:  reflect.ValueOf(false),
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

			if got := encodeBool(tt.arg); got != tt.want {
				t.Errorf("encodeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
