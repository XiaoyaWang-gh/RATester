package proxy

import (
	"fmt"
	"net"
	"testing"
)

func TesthandleConn_1(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		s    *Sock5ModeServer
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

			tt.s.handleConn(tt.args.c)
		})
	}
}
