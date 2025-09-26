package transform

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestToFormatMark_1(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		want    metadecoders.Format
		wantErr bool
	}{
		{
			name:    "Test Case 1",
			format:  "json",
			want:    "json",
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			format:  "yaml",
			want:    "yaml",
			wantErr: false,
		},
		{
			name:    "Test Case 3",
			format:  "toml",
			want:    "toml",
			wantErr: false,
		},
		{
			name:    "Test Case 4",
			format:  "unknown",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := toFormatMark(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("toFormatMark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("toFormatMark() = %v, want %v", got, tt.want)
			}
		})
	}
}
