package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertInt8_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid int8",
			value: "123",
			want:  reflect.ValueOf(int8(123)),
		},
		{
			name:  "invalid int8",
			value: "256",
			want:  reflect.ValueOf(int8(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertInt8(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
