package compress

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/andybalholm/brotli"
)

func TestNewCompressionWriter_1(t *testing.T) {
	type args struct {
		algo string
		in   io.Writer
	}
	tests := []struct {
		name    string
		args    args
		want    *compressionWriter
		wantErr bool
	}{
		{
			name: "Test with brotli",
			args: args{
				algo: brotliName,
				in:   &bytes.Buffer{},
			},
			want: &compressionWriter{
				compression: &brotli.Writer{},
				alg:         brotliName,
			},
			wantErr: false,
		},
		{
			name: "Test with unknown algo",
			args: args{
				algo: "unknown",
				in:   &bytes.Buffer{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newCompressionWriter(tt.args.algo, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCompressionWriter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCompressionWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
