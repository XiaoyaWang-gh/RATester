package tcp

import (
	"fmt"
	"testing"
)

func TestIsASCII_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"ASCII string", "hello", true},
		{"Non-ASCII string", "😀", false},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isASCII(tt.arg); got != tt.want {
				t.Errorf("isASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
