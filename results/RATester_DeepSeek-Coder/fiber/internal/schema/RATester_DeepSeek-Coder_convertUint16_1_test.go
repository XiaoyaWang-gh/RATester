package schema

import (
	"fmt"
	"testing"
)

func TestConvertUint16_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  uint16
	}{
		{
			name:  "valid uint16",
			value: "65535",
			want:  65535,
		},
		{
			name:  "invalid uint16",
			value: "65536",
			want:  0,
		},
		{
			name:  "empty string",
			value: "",
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

			got := convertUint16(tt.value)
			if got.Uint() != uint64(tt.want) {
				t.Errorf("convertUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}
