package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertFloat64_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid float",
			value: "123.456",
			want:  reflect.ValueOf(123.456),
		},
		{
			name:  "invalid float",
			value: "abc",
			want:  reflect.ValueOf(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertFloat64(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
