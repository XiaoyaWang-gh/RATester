package client

import (
	"fmt"
	"testing"
)

func TestNewUdpConn_2(t *testing.T) {
	type args struct {
		localAddr   string
		rAddr       string
		md5Password string
	}
	tests := []struct {
		name string
		s    *TRPClient
		args args
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

			tt.s.newUdpConn(tt.args.localAddr, tt.args.rAddr, tt.args.md5Password)
		})
	}
}
