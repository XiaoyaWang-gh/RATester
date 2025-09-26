package template

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestHTMLEscape_1(t *testing.T) {
	type args struct {
		w io.Writer
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with simple string",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<test>"),
			},
			want: "&lt;test&gt;",
		},
		{
			name: "Test with special characters",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("&<>\"'"),
			},
			want: "&amp;&lt;&gt;&quot;&#39;",
		},
		{
			name: "Test with empty string",
			args: args{
				w: &bytes.Buffer{},
				b: []byte(""),
			},
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

			HTMLEscape(tt.args.w, tt.args.b)
			got := tt.args.w.(*bytes.Buffer).String()
			if got != tt.want {
				t.Errorf("HTMLEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
