package helpers

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReaderToBytes_1(t *testing.T) {
	type args struct {
		lines io.Reader
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test case 1",
			args: args{
				lines: strings.NewReader("Hello, World!"),
			},
			want: []byte("Hello, World!"),
		},
		{
			name: "Test case 2",
			args: args{
				lines: strings.NewReader(""),
			},
			want: []byte(""),
		},
		{
			name: "Test case 3",
			args: args{
				lines: strings.NewReader("1234567890"),
			},
			want: []byte("1234567890"),
		},
		{
			name: "Test case 4",
			args: args{
				lines: strings.NewReader("abcdefghijklmnopqrstuvwxyz"),
			},
			want: []byte("abcdefghijklmnopqrstuvwxyz"),
		},
		{
			name: "Test case 5",
			args: args{
				lines: strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			},
			want: []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ReaderToBytes(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReaderToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
