package template

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestJSEscape_1(t *testing.T) {
	type args struct {
		w io.Writer
		b []byte
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "Test case 1",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<script>alert('XSS')</script>"),
			},
			wantW: "&lt;script&gt;alert('XSS')&lt;/script&gt;",
		},
		{
			name: "Test case 2",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<img src=x onerror=alert('XSS')>"),
			},
			wantW: "&lt;img src=x onerror=alert('XSS')&gt;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			JSEscape(tt.args.w, tt.args.b)
			if gotW := tt.args.w.(*bytes.Buffer).String(); gotW != tt.wantW {
				t.Errorf("JSEscape() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
