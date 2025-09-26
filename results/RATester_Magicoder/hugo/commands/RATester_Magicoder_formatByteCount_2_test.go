package commands

import (
	"fmt"
	"testing"
)

func TestformatByteCount_2(t *testing.T) {
	tests := []struct {
		name string
		b    uint64
		want string
	}{
		{"1000 bytes", 1000, "1000 B"},
		{"1024 bytes", 1024, "1.0 KB"},
		{"1500 bytes", 1500, "1.5 KB"},
		{"1000000 bytes", 1000000, "1.0 MB"},
		{"1000000000 bytes", 1000000000, "1.0 GB"},
		{"1000000000000 bytes", 1000000000000, "1.0 TB"},
		{"1000000000000000 bytes", 1000000000000000, "1.0 PB"},
		{"1000000000000000000 bytes", 1000000000000000000, "1.0 EB"},
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
