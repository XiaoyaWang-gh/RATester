package http

import (
	"fmt"
	"testing"
)

func TestIsASCII_2(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"empty string", "", true},
		{"ascii string", "hello, world", true},
		{"non-ascii string", "日本語", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsASCII(tt.arg); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
