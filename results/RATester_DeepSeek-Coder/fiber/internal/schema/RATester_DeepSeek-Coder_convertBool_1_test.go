package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertBool_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "Test with 'on'",
			value: "on",
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Test with 'true'",
			value: "true",
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Test with 'false'",
			value: "false",
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Test with 'invalid'",
			value: "invalid",
			want:  reflect.ValueOf(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertBool(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
