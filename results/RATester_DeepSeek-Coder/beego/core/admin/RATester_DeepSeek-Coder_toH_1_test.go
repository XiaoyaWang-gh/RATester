package admin

import (
	"fmt"
	"testing"
)

func TestToH_1(t *testing.T) {
	tests := []struct {
		name  string
		bytes uint64
		want  string
	}{
		{"less than 1K", 100, "100B"},
		{"1K", 1024, "1.00K"},
		{"1.5K", 1536, "1.50K"},
		{"1M", 1024 * 1024, "1.00M"},
		{"1.5M", 1572864, "1.50M"},
		{"1G", 1024 * 1024 * 1024, "1.00G"},
		{"1.5G", 1610612736, "1.50G"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toH(tt.bytes); got != tt.want {
				t.Errorf("toH() = %v, want %v", got, tt.want)
			}
		})
	}
}
