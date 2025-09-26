package common

import (
	"fmt"
	"io"
	"testing"
)

func TestCopyBuffer_1(t *testing.T) {
	type args struct {
		dst   io.Writer
		src   io.Reader
		label []string
	}
	tests := []struct {
		name        string
		args        args
		wantWritten int64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotWritten, err := CopyBuffer(tt.args.dst, tt.args.src, tt.args.label...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyBuffer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWritten != tt.wantWritten {
				t.Errorf("CopyBuffer() = %v, want %v", gotWritten, tt.wantWritten)
			}
		})
	}
}
