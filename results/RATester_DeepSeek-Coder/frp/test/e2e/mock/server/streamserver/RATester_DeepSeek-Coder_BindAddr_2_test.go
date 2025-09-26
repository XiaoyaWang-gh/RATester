package streamserver

import (
	"fmt"
	"net"
	"testing"
)

func TestBindAddr_2(t *testing.T) {
	type fields struct {
		netType     Type
		bindAddr    string
		bindPort    int
		respContent []byte
		handler     func(net.Conn)
		l           net.Listener
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestBindAddr",
			fields: fields{
				netType:  "tcp",
				bindAddr: "127.0.0.1",
			},
			want: "127.0.0.1",
		},
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
			if got := s.BindAddr(); got != tt.want {
				t.Errorf("Server.BindAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
