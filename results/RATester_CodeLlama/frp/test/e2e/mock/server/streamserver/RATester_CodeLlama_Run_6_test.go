package streamserver

import (
	"fmt"
	"net"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestRun_6(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		netType     Type
		bindAddr    string
		bindPort    int
		respContent []byte
		handler     func(net.Conn)
		l           net.Listener
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				netType:     "tcp",
				bindAddr:    "127.0.0.1",
				bindPort:    8080,
				respContent: []byte("hello world"),
				handler: func(c net.Conn) {
					c.Write([]byte("hello world"))
					c.Close()
				},
			},
			wantErr: false,
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
			if err := s.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
