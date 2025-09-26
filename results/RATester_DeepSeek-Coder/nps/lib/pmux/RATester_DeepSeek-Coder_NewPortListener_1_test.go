package pmux

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewPortListener_1(t *testing.T) {
	type args struct {
		connCh chan *PortConn
		addr   net.Addr
	}

	tests := []struct {
		name string
		args args
		want *PortListener
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

			if got := NewPortListener(tt.args.connCh, tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
