package hugio

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestToWriteCloser_1(t *testing.T) {
	type args struct {
		w io.Writer
	}
	tests := []struct {
		name string
		args args
		want io.WriteCloser
	}{
		{
			name: "test case 1",
			args: args{
				w: nil,
			},
			want: nil,
		},
		{
			name: "test case 2",
			args: args{
				w: &bytes.Buffer{},
			},
			want: &struct {
				io.Writer
				io.Closer
			}{
				Writer: &bytes.Buffer{},
				Closer: io.NopCloser(nil),
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

			if got := ToWriteCloser(tt.args.w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToWriteCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
