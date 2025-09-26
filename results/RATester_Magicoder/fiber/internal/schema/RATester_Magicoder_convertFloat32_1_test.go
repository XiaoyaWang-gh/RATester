package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertFloat32_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid float",
			value: "3.14",
			want:  reflect.ValueOf(float32(3.14)),
		},
		{
			name:  "invalid float",
			value: "invalid",
			want:  reflect.ValueOf(float32(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertFloat32(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
