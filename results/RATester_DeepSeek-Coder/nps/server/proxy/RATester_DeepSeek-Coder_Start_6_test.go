package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestStart_6(t *testing.T) {
	type fields struct {
		BaseServer BaseServer
		listener   net.Listener
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
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

			s := &Sock5ModeServer{
				BaseServer: tt.fields.BaseServer,
				listener:   tt.fields.listener,
			}
			if err := s.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Sock5ModeServer.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
