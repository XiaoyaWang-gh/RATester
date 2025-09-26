package streamserver

import (
	"fmt"
	"net"
	"testing"
)

func TestHandle_4(t *testing.T) {
	type fields struct {
		netType     Type
		bindAddr    string
		bindPort    int
		respContent []byte
		handler     func(net.Conn)
		l           net.Listener
	}
	type args struct {
		c net.Conn
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

			s := &Server{
				netType:     tt.fields.netType,
				bindAddr:    tt.fields.bindAddr,
				bindPort:    tt.fields.bindPort,
				respContent: tt.fields.respContent,
				handler:     tt.fields.handler,
				l:           tt.fields.l,
			}
			s.handle(tt.args.c)
		})
	}
}
