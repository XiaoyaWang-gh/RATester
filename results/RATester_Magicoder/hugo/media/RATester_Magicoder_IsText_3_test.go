package media

import (
	"fmt"
	"testing"
)

func TestIsText_3(t *testing.T) {
	tests := []struct {
		name string
		m    Type
		want bool
	}{
		{
			name: "text",
			m:    Type{MainType: "text"},
			want: true,
		},
		{
			name: "javascript",
			m:    Type{SubType: "javascript"},
			want: true,
		},
		{
			name: "json",
			m:    Type{SubType: "json"},
			want: true,
		},
		{
			name: "rss",
			m:    Type{SubType: "rss"},
			want: true,
		},
		{
			name: "xml",
			m:    Type{SubType: "xml"},
			want: true,
		},
		{
			name: "svg",
			m:    Type{SubType: "svg"},
			want: true,
		},
		{
			name: "toml",
			m:    Type{SubType: "toml"},
			want: true,
		},
		{
			name: "yml",
			m:    Type{SubType: "yml"},
			want: true,
		},
		{
			name: "yaml",
			m:    Type{SubType: "yaml"},
			want: true,
		},
		{
			name: "other",
			m:    Type{SubType: "other"},
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

			if got := tt.m.IsText(); got != tt.want {
				t.Errorf("IsText() = %v, want %v", got, tt.want)
			}
		})
	}
}
