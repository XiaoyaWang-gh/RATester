package common

import (
	"fmt"
	"testing"
)

func TestFormatAddress_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test1",
			s:    "127.0.0.1:8080",
			want: "127.0.0.1:8080",
		},
		{
			name: "test2",
			s:    "8080",
			want: "127.0.0.1:8080",
		},
		{
			name: "test3",
			s:    "127.0.0.1",
			want: "127.0.0.1:127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatAddress(tt.s); got != tt.want {
				t.Errorf("FormatAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
