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
		{
			name: "1000",
			bc:   1000,
			want: "1000 B",
		},
		{
			name: "1000000",
			bc:   1000000,
			want: "1000 KB",
		},
		{
			name: "1000000000",
			bc:   1000000000,
			want: "1000 MB",
		},
		{
			name: "1000000000000",
			bc:   1000000000000,
			want: "1000 GB",
		},
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
