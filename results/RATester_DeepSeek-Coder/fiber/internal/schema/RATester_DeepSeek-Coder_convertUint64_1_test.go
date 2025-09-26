package schema

import (
	"fmt"
	"testing"
)

func TestConvertUint64_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		value string
		want  uint64
	}{
		{
			name:  "valid uint64",
			value: "18446744073709551615",
			want:  18446744073709551615,
		},
		{
			name:  "invalid uint64",
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

			got := convertUint64(tt.value)
			if got.Uint() != tt.want {
				t.Errorf("convertUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
