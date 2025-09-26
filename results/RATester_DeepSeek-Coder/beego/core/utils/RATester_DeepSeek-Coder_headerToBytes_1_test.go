package utils

import (
	"fmt"
	"io"
	"net/textproto"
	"testing"
)

func TestHeaderToBytes_1(t *testing.T) {
	type args struct {
		w io.Writer
		t textproto.MIMEHeader
	}
	tests := []struct {
		name    string
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

			if err := headerToBytes(tt.args.w, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("headerToBytes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
