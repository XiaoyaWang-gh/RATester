package common

import (
	"fmt"
	"io"
	"reflect"
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
		wantErr     error
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

			gotWritten, gotErr := CopyBuffer(tt.args.dst, tt.args.src, tt.args.label...)
			if !reflect.DeepEqual(gotWritten, tt.wantWritten) {
				t.Errorf("CopyBuffer() gotWritten = %v, want %v", gotWritten, tt.wantWritten)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("CopyBuffer() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
