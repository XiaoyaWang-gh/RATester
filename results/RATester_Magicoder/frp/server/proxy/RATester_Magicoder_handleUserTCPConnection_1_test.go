package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestHandleUserTCPConnection_1(t *testing.T) {
	type args struct {
		userConn net.Conn
	}
	tests := []struct {
		name string
		pxy  *BaseProxy
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

			tt.pxy.handleUserTCPConnection(tt.args.userConn)
		})
	}
}
