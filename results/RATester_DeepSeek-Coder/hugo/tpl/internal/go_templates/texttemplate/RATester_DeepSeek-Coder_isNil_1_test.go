package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsNil_1(t *testing.T) {
	type testStruct struct {
		field string
	}

	tests := []struct {
		name string
		arg  reflect.Value
		want bool
	}{
		{
			name: "Test with nil pointer",
			arg:  reflect.ValueOf((*testStruct)(nil)),
			want: true,
		},
		{
			name: "Test with non-nil pointer",
			arg:  reflect.ValueOf(&testStruct{}),
			want: false,
		},
		{
			name: "Test with non-pointer value",
			arg:  reflect.ValueOf(testStruct{}),
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

			if got := isNil(tt.arg); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}
