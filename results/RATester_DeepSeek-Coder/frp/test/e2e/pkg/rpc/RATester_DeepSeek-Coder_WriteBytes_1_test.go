package rpc

import (
	"fmt"
	"io"
	"testing"
)

func TestWriteBytes_1(t *testing.T) {
	type args struct {
		w   io.Writer
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
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

			gotN, err := WriteBytes(tt.args.w, tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("WriteBytes() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
