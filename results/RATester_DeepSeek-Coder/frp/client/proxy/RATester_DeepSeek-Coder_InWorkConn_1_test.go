package proxy

import (
	"fmt"
	"net"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestInWorkConn_1(t *testing.T) {
	type args struct {
		conn             net.Conn
		startWorkConnMsg *msg.StartWorkConn
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

			pxy := &XTCPProxy{}
			pxy.InWorkConn(tt.args.conn, tt.args.startWorkConnMsg)
		})
	}
}
