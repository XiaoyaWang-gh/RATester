package helpers

import (
	"fmt"
	"testing"
)

func TestFormatByteCount_1(t *testing.T) {
	tests := []struct {
		name string
		bc   uint64
		want string
	}{
		{"Test 1GB", 1024 * 1024 * 1024, "1.00 GB"},
		{"Test 1.5GB", 1536 * 1024 * 1024, "1.50 GB"},
		{"Test 1MB", 1024 * 1024, "1.00 MB"},
		{"Test 1.5MB", 1536 * 1024, "1.50 MB"},
		{"Test 1KB", 1024, "1.00 KB"},
		{"Test 1.5KB", 1536, "1.50 KB"},
		{"Test 1B", 1, "1 B"},
		{"Test 0B", 0, "0 B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatByteCount(tt.bc); got != tt.want {
				t.Errorf("FormatByteCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
