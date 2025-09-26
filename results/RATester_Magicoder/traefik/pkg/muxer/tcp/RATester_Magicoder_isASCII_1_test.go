package tcp

import (
	"fmt"
	"testing"
)

func TestIsASCII_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "ASCII string",
			s:    "Hello, World!",
			want: true,
		},
		{
			name: "Non-ASCII string",
			s:    "你好，世界!",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isASCII(tt.s); got != tt.want {
				t.Errorf("isASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
