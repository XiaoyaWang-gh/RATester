package helpers

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReaderToString_1(t *testing.T) {
	type args struct {
		lines io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				lines: strings.NewReader("Hello, world!"),
			},
			want: "Hello, world!\n",
		},
		{
			name: "Test case 2",
			args: args{
				lines: strings.NewReader(""),
			},
			want: "",
		},
		{
			name: "Test case 3",
			args: args{
				lines: strings.NewReader("Golang is fun!"),
			},
			want: "Golang is fun!\n",
		},
		{
			name: "Test case 4",
			args: args{
				lines: nil,
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

			if got := ReaderToString(tt.args.lines); got != tt.want {
				t.Errorf("ReaderToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
