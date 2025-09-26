package utils

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func Testbase64Wrap_1(t *testing.T) {
	type args struct {
		w io.Writer
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				w: bytes.NewBuffer([]byte{}),
				b: []byte("Hello, World!"),
			},
			want: []byte("SGVsbG8sIFdvcmxkIQ==\r\n"),
		},
		{
			name: "Test Case 2",
			args: args{
				w: bytes.NewBuffer([]byte{}),
				b: []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit."),
			},
			want: []byte("TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=\r\n"),
		},
		{
			name: "Test Case 3",
			args: args{
				w: bytes.NewBuffer([]byte{}),
				b: []byte(""),
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			base64Wrap(tt.args.w, tt.args.b)
			got := tt.args.w.(*bytes.Buffer).Bytes()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base64Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
