package text

import (
	"fmt"
	"testing"
)

func TestChomp_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "string without newline",
			s:    "Hello, World!",
			want: "Hello, World!",
		},
		{
			name: "string with newline",
			s:    "Hello, World!\n",
			want: "Hello, World!",
		},
		{
			name: "string with carriage return",
			s:    "Hello, World!\r",
			want: "Hello, World!",
		},
		{
			name: "string with newline and carriage return",
			s:    "Hello, World!\n\r",
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Chomp(tt.s); got != tt.want {
				t.Errorf("Chomp() = %v, want %v", got, tt.want)
			}
		})
	}
}
