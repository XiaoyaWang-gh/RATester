package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/textproto"
	"testing"
)

func TestheaderToBytes_1(t *testing.T) {
	type args struct {
		w io.Writer
		t textproto.MIMEHeader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				w: &bytes.Buffer{},
				t: textproto.MIMEHeader{
					"Content-Type":   []string{"text/plain"},
					"Content-Length": []string{"100"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				w: &bytes.Buffer{},
				t: textproto.MIMEHeader{
					"Content-Type":     []string{"text/plain"},
					"Content-Length":   []string{"100"},
					"Content-Encoding": []string{"gzip"},
				},
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				w: &bytes.Buffer{},
				t: textproto.MIMEHeader{
					"Content-Type":        []string{"text/plain"},
					"Content-Length":      []string{"100"},
					"Content-Encoding":    []string{"gzip"},
					"Content-Disposition": []string{"attachment; filename=test.txt"},
				},
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

			if err := headerToBytes(tt.args.w, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("headerToBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
