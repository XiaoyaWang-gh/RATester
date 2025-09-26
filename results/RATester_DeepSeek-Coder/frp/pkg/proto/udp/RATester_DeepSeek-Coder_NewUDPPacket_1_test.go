package udp

import (
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestNewUDPPacket_1(t *testing.T) {
	type args struct {
		buf   []byte
		laddr *net.UDPAddr
		raddr *net.UDPAddr
	}
	tests := []struct {
		name string
		args args
		want *msg.UDPPacket
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

			got := NewUDPPacket(tt.args.buf, tt.args.laddr, tt.args.raddr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUDPPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}
