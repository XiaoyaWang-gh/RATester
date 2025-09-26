package schema

import (
	"fmt"
	"testing"
)

func TestConvertInt_1(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int
	}{
		{
			name:  "Test with valid integer",
			value: "10",
			want:  10,
		},
		{
			name:  "Test with invalid integer",
			value: "invalid",
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

			got := convertInt(tt.value)
			if got.Int() != int64(tt.want) {
				t.Errorf("convertInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
