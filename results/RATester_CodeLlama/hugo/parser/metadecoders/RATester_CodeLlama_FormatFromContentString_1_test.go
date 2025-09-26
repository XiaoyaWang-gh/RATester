package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromContentString_1(t *testing.T) {
	tests := []struct {
		name string
		d    Decoder
		data string
		want Format
	}{
		{
			name: "csv",
			d:    Decoder{Delimiter: ','},
			data: "a,b,c",
			want: CSV,
		},
		{
			name: "json",
			d:    Decoder{},
			data: "{\"a\":\"b\"}",
			want: JSON,
		},
		{
			name: "yaml",
			d:    Decoder{},
			data: "a: b",
			want: YAML,
		},
		{
			name: "xml",
			d:    Decoder{},
			data: "<a>b</a>",
			want: XML,
		},
		{
			name: "toml",
			d:    Decoder{},
			data: "a = b",
			want: TOML,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.FormatFromContentString(tt.data); got != tt.want {
				t.Errorf("Decoder.FormatFromContentString() = %v, want %v", got, tt.want)
			}
		})
	}
}
