package server

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestHandleConnection_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		conn     net.Conn
		internal bool
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

			svr := &Service{}
			svr.handleConnection(tt.args.ctx, tt.args.conn, tt.args.internal)
		})
	}
}
