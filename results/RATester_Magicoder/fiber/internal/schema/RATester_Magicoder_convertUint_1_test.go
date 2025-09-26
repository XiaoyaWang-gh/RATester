package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertUint_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid uint",
			value: "123",
			want:  reflect.ValueOf(uint(123)),
		},
		{
			name:  "invalid uint",
			value: "abc",
			want:  reflect.ValueOf(uint(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertUint(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUint() = %v, want %v", got, tt.want)
			}
		})
	}
}
