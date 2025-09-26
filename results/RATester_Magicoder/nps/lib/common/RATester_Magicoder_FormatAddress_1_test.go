package common

import (
	"fmt"
	"testing"
)

func TestFormatAddress_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Test with valid IP", "8080", "127.0.0.1:8080"},
		{"Test with valid IP and port", "127.0.0.1:8080", "127.0.0.1:8080"},
		{"Test with invalid IP", "invalid", "127.0.0.1:invalid"},
		{"Test with empty string", "", "127.0.0.1:"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatAddress(tt.arg); got != tt.want {
				t.Errorf("FormatAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
