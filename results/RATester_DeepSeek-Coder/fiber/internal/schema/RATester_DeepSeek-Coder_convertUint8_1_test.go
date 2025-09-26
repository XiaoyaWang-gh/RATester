package schema

import (
	"fmt"
	"testing"
)

func TestConvertUint8_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  uint8
	}{
		{
			name:  "valid uint8",
			value: "123",
			want:  123,
		},
		{
			name:  "invalid uint8",
			value: "256",
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

			got := convertUint8(tt.value)
			if got.Uint() != uint64(tt.want) {
				t.Errorf("convertUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
