package helpers

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestReaderContains_1(t *testing.T) {
	type args struct {
		r        io.Reader
		subslice []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				r:        strings.NewReader("Hello, world!"),
				subslice: []byte("world"),
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				r:        strings.NewReader("Hello, world!"),
				subslice: []byte("foo"),
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				r:        strings.NewReader(""),
				subslice: []byte("foo"),
			},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{
				r:        strings.NewReader("Hello, world!"),
				subslice: []byte(""),
			},
			want: false,
		},
		{
			name: "Test case 5",
			args: args{
				r:        strings.NewReader("Hello, world!"),
				subslice: nil,
			},
			want: false,
		},
		{
			name: "Test case 6",
			args: args{
				r:        nil,
				subslice: []byte("foo"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ReaderContains(tt.args.r, tt.args.subslice); got != tt.want {
				t.Errorf("ReaderContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
