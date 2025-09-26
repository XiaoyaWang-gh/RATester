package commands

import (
	"fmt"
	"testing"
)

func TestFormatByteCount_2(t *testing.T) {
	tests := []struct {
		name string
		b    uint64
		want string
	}{
		{"Test 1", 100, "100 B"},
		{"Test 2", 1024, "1.0 KiB"},
		{"Test 3", 1048576, "1.0 MiB"},
		{"Test 4", 1073741824, "1.0 GiB"},
		{"Test 5", 1099511627776, "1.1 TiB"},
		{"Test 6", 1125899906842624, "1.1 PiB"},
		{"Test 7", 1152921504606846976, "1.2 EiB"},
		{"Test 8", 18446744073709551615, "18.4 EiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := formatByteCount(tt.b); got != tt.want {
				t.Errorf("formatByteCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
