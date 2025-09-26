package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromString_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want Format
	}{
		{
			name: "Test YAML",
			arg:  "yaml",
			want: YAML,
		},
		{
			name: "Test YML",
			arg:  "yml",
			want: YAML,
		},
		{
			name: "Test JSON",
			arg:  "json",
			want: JSON,
		},
		{
			name: "Test TOML",
			arg:  "toml",
			want: TOML,
		},
		{
			name: "Test ORG",
			arg:  "org",
			want: "org",
		},
		{
			name: "Test CSV",
			arg:  "csv",
			want: "csv",
		},
		{
			name: "Test XML",
			arg:  "xml",
			want: "xml",
		},
		{
			name: "Test Unknown",
			arg:  "unknown",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatFromString(tt.arg); got != tt.want {
				t.Errorf("FormatFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
