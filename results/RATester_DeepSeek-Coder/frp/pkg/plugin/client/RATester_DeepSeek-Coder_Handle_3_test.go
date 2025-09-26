package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"
)

func TestHandle_3(t *testing.T) {
	type args struct {
		ctx   context.Context
		conn  io.ReadWriteCloser
		local net.Conn
		extra *ExtraInfo
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

			uds := &UnixDomainSocketPlugin{}
			uds.Handle(tt.args.ctx, tt.args.conn, tt.args.local, tt.args.extra)
		})
	}
}
