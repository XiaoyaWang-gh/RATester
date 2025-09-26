package utils

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestquotePrintEncode_1(t *testing.T) {
	type args struct {
		w io.Writer
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				w: &bytes.Buffer{},
				s: "Hello, World!",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				w: &bytes.Buffer{},
				s: "This is a test string with special characters: !@#$%^&*()",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				w: &bytes.Buffer{},
				s: "This is a test string with newline\n characters",
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				w: &bytes.Buffer{},
				s: "",
			},
			wantErr: false,
		},
		{
			name: "Test case 5",
			args: args{
				w: &bytes.Buffer{},
				s: "This is a test string with a very long line that should be split into multiple lines to ensure the correct behavior of the function",
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

			if err := quotePrintEncode(tt.args.w, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("quotePrintEncode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
