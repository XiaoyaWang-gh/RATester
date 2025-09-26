package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromString_1(t *testing.T) {
	tests := []struct {
		name      string
		formatStr string
		want      Format
	}{
		{
			name:      "Test YAML",
			formatStr: "yaml",
			want:      YAML,
		},
		{
			name:      "Test YML",
			formatStr: "yml",
			want:      YAML,
		},
		{
			name:      "Test JSON",
			formatStr: "json",
			want:      JSON,
		},
		{
			name:      "Test TOML",
			formatStr: "toml",
			want:      TOML,
		},
		{
			name:      "Test ORG",
			formatStr: "org",
			want:      ORG,
		},
		{
			name:      "Test CSV",
			formatStr: "csv",
			want:      CSV,
		},
		{
			name:      "Test XML",
			formatStr: "xml",
			want:      XML,
		},
		{
			name:      "Test Unsupported Format",
			formatStr: "unsupported",
			want:      "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FormatFromString(tt.formatStr); got != tt.want {
				t.Errorf("FormatFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
