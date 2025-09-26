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
		{
			name: "ASCII string",
			arg:  "Hello, World!",
			want: true,
		},
		{
			name: "Non-ASCII string",
			arg:  "你好，世界!",
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

			if got := IsASCII(tt.arg); got != tt.want {
				t.Errorf("IsASCII() = %v, want %v", got, tt.want)
			}
		})
	}
}
