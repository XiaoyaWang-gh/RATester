package client

import (
	"fmt"
	"net"
	"testing"
)

func TesthandleUdp_1(t *testing.T) {
	type args struct {
		serverConn net.Conn
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

			tt.s.handleUdp(tt.args.serverConn)
		})
	}
}
