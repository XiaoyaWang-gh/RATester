package server

import (
	"fmt"
	"syscall"
	"testing"
)

func TestControlReusePort_1(t *testing.T) {
	type args struct {
		network string
		address string
		c       syscall.RawConn
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

			err := controlReusePort(tt.args.network, tt.args.address, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("controlReusePort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
