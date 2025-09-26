package urls

import (
	"fmt"
	"testing"
)

func TestString_23(t *testing.T) {
	tests := []struct {
		name string
		b    BaseURL
		want string
	}{
		{
			name: "Test case 1",
			b:    BaseURL{WithPath: "http://example.com"},
			want: "http://example.com",
		},
		{
			name: "Test case 2",
			b:    BaseURL{WithPath: "https://example.com"},
			want: "https://example.com",
		},
		{
			name: "Test case 3",
			b:    BaseURL{WithPath: "ftp://example.com"},
			want: "ftp://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.b.String(); got != tt.want {
				t.Errorf("BaseURL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
