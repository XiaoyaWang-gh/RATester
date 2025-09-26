package nathole

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestSendSidMessageToRandomPorts_1(t *testing.T) {
	type args struct {
		ctx      context.Context
		conn     *net.UDPConn
		addrs    []string
		count    int
		sendFunc func(*net.UDPConn, string) error
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

			sendSidMessageToRandomPorts(tt.args.ctx, tt.args.conn, tt.args.addrs, tt.args.count, tt.args.sendFunc)
		})
	}
}
