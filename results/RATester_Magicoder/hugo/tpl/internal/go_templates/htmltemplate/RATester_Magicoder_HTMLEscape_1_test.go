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
	}{
		{
			name: "Test Case 1",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<script>alert('Hello, World!')</script>"),
			},
		},
		{
			name: "Test Case 2",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<a href='javascript:alert(1);'>Click me</a>"),
			},
		},
		{
			name: "Test Case 3",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("<img src='x' onerror='alert(1);'>"),
			},
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
			// Add assertions here to validate the output of HTMLEscape
		})
	}
}
