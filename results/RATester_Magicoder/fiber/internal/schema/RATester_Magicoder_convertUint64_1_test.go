package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertUint64_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid uint64",
			value: "1234567890",
			want:  reflect.ValueOf(uint64(1234567890)),
		},
		{
			name:  "invalid uint64",
			value: "invalid",
			want:  reflect.ValueOf(uint64(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertUint64(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
