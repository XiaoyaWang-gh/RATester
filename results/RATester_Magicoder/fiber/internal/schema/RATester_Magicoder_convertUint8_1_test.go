package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertUint8_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid uint8",
			value: "123",
			want:  reflect.ValueOf(uint8(123)),
		},
		{
			name:  "invalid uint8",
			value: "256",
			want:  reflect.ValueOf(uint8(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertUint8(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
