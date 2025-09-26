package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestAuth_1(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name    string
		s       *Sock5ModeServer
		args    args
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

			if err := tt.s.Auth(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Sock5ModeServer.Auth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
