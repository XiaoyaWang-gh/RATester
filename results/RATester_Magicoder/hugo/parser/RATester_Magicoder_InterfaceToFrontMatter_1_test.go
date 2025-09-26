package parser

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestInterfaceToFrontMatter_1(t *testing.T) {
	type args struct {
		in     any
		format metadecoders.Format
		w      io.Writer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				in:     nil,
				format: metadecoders.YAML,
				w:      &bytes.Buffer{},
			},
			wantErr: true,
		},
		{
			name: "Test case 2",
			args: args{
				in:     map[string]interface{}{"key": "value"},
				format: metadecoders.YAML,
				w:      &bytes.Buffer{},
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				in:     map[string]interface{}{"key": "value"},
				format: metadecoders.TOML,
				w:      &bytes.Buffer{},
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				in:     map[string]interface{}{"key": "value"},
				format: "unknown",
				w:      &bytes.Buffer{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := InterfaceToFrontMatter(tt.args.in, tt.args.format, tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToFrontMatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
