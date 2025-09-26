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
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("test"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				w: &bytes.Buffer{},
				b: []byte(""),
			},
		},
		{
			name: "Test case 3",
			args: args{
				w: &bytes.Buffer{},
				b: []byte("test with special characters: !@#$%^&*()"),
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

			JSEscape(tt.args.w, tt.args.b)
			// Add assertions here to check the output of JSEscape
		})
	}
}
