package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertInt32_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid int32",
			value: "123",
			want:  reflect.ValueOf(int32(123)),
		},
		{
			name:  "invalid int32",
			value: "abc",
			want:  reflect.ValueOf(int32(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertInt32(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
