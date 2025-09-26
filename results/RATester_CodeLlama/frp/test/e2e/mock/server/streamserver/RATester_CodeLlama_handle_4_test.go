package streamserver

import (
	"fmt"
	"net"
	"testing"
)

func TestHandle_4(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		s    *Server
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

			tt.s.handle(tt.args.c)
		})
	}
}
