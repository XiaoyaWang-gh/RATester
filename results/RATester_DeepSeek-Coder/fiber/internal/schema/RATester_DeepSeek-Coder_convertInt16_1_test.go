package schema

import (
	"fmt"
	"testing"
)

func TestConvertInt16_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  int16
	}{
		{
			name:  "valid int16",
			value: "32767",
			want:  32767,
		},
		{
			name:  "invalid int16",
			value: "32768",
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

			got := convertInt16(tt.value)
			if got.Int() != int64(tt.want) {
				t.Errorf("convertInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}
