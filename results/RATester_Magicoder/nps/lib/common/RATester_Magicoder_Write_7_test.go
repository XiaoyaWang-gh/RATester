package common

import (
	"fmt"
	"io"
	"testing"
)

func TestWrite_7(t *testing.T) {
	type args struct {
		w io.Writer
	}
	tests := []struct {
		name    string
		h       *UDPHeader
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

			if err := tt.h.Write(tt.args.w); (err != nil) != tt.wantErr {
				t.Errorf("UDPHeader.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
