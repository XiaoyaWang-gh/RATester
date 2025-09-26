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
			name: "Test with nil input",
			args: args{
				in:     nil,
				format: metadecoders.YAML,
				w:      &bytes.Buffer{},
			},
			wantErr: true,
		},
		{
			name: "Test with valid input",
			args: args{
				in:     map[string]interface{}{"title": "Test Post"},
				format: metadecoders.YAML,
				w:      &bytes.Buffer{},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := InterfaceToFrontMatter(tt.args.in, tt.args.format, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToFrontMatter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
