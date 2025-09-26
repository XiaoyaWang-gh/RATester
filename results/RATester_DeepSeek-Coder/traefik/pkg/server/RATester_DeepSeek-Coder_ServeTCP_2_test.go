package server

import (
	"fmt"
	"net"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tcp"
)

func TestServeTCP_2(t *testing.T) {
	type fields struct {
		Listener net.Listener
		connChan chan net.Conn
		errChan  chan error
	}
	type args struct {
		conn tcp.WriteCloser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			h := &httpForwarder{
				Listener: tt.fields.Listener,
				connChan: tt.fields.connChan,
				errChan:  tt.fields.errChan,
			}
			h.ServeTCP(tt.args.conn)
		})
	}
}
