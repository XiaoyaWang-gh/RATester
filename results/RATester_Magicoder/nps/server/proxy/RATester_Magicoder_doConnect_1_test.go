package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TestdoConnect_1(t *testing.T) {
	type args struct {
		c       net.Conn
		command uint8
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

			s := &Sock5ModeServer{}
			s.doConnect(tt.args.c, tt.args.command)
		})
	}
}
