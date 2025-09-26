package echo

import (
	"fmt"
	"io"
	"testing"
)

func TestStream_1(t *testing.T) {
	type args struct {
		code        int
		contentType string
		r           io.Reader
	}
	tests := []struct {
		name    string
		c       *context
		args    args
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

			if err := tt.c.Stream(tt.args.code, tt.args.contentType, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("context.Stream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
