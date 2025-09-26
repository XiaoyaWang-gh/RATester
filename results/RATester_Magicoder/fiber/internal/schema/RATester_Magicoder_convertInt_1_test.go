package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconvertInt_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  reflect.Value
	}{
		{
			name:  "valid int",
			value: "123",
			want:  reflect.ValueOf(123),
		},
		{
			name:  "invalid int",
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

			if got := convertInt(tt.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
