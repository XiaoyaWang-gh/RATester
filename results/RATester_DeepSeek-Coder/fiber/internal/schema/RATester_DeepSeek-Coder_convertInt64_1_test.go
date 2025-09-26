package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertInt64_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid int",
			value: "1234567890",
			want:  reflect.ValueOf(int64(1234567890)),
		},
		{
			name:  "invalid int",
			value: "invalid",
			want:  invalidValue,
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
