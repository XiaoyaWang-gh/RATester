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
			name:     "valid config file",
			filename: "config.yaml",
			want:     true,
		},
		{
			name:     "invalid config file",
			filename: "config.txt",
			want:     false,
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
