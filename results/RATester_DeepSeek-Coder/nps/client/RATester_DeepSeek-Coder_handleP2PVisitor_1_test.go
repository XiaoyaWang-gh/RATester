package client

import (
	"fmt"
	"net"
	"testing"

	"ehang.io/nps/lib/config"
)

func TestHandleP2PVisitor_1(t *testing.T) {
	type args struct {
		localTcpConn net.Conn
		config       *config.CommonConfig
		l            *config.LocalServer
	}
	tests := []struct {
		name string
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

			handleP2PVisitor(tt.args.localTcpConn, tt.args.config, tt.args.l)
		})
	}
}
