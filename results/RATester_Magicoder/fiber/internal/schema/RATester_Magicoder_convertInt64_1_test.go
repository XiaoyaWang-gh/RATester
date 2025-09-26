package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertInt64_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid int64",
			value: "1234567890",
			want:  reflect.ValueOf(int64(1234567890)),
		},
		{
			name:  "invalid int64",
			value: "invalid",
			want:  reflect.ValueOf(int64(0)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertInt64(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
