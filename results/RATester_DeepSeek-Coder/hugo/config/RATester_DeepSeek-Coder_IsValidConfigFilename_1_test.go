package config

import (
	"fmt"
	"testing"
)

func TestIsValidConfigFilename_1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "Test with valid config file",
			filename: "config.json",
			want:     true,
		},
		{
			name:     "Test with invalid config file",
			filename: "config.txt",
			want:     false,
		},
		{
			name:     "Test with no extension",
			filename: "config",
			want:     false,
		},
		{
			name:     "Test with uppercase extension",
			filename: "config.JSON",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsValidConfigFilename(tt.filename); got != tt.want {
				t.Errorf("IsValidConfigFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
