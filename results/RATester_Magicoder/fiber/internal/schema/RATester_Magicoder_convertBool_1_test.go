package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertBool_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "Test case 1",
			value: "on",
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Test case 2",
			value: "true",
			want:  reflect.ValueOf(true),
		},
		{
			name:  "Test case 3",
			value: "false",
			want:  reflect.ValueOf(false),
		},
		{
			name:  "Test case 4",
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
