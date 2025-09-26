package ssh

import (
	"fmt"
	"net"
	"testing"
)

func TestHandleConn_1(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		g    *Gateway
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

			g := &Gateway{}
			g.handleConn(tt.args.conn)
		})
	}
}
