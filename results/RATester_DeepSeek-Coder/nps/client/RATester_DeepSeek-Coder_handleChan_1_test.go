package client

import (
	"fmt"
	"net"
	"testing"
)

func TestHandleChan_1(t *testing.T) {
	type args struct {
		src net.Conn
	}
	tests := []struct {
		name string
		s    *TRPClient
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

			tt.s.handleChan(tt.args.src)
		})
	}
}
