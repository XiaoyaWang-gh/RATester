package schema

import (
	"fmt"
	"testing"
)

func TestConvertInt8_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  int8
	}{
		{
			name:  "valid int8",
			value: "127",
			want:  127,
		},
		{
			name:  "negative int8",
			value: "-128",
			want:  -128,
		},
		{
			name:  "out of range int8",
			value: "128",
			want:  0,
		},
		{
			name:  "not a number",
			value: "abc",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := convertInt8(tt.value)
			if got.Int() != int64(tt.want) {
				t.Errorf("convertInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
