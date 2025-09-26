package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewFakeUDPConn_1(t *testing.T) {
	type args struct {
		l     *UDPListener
		laddr net.Addr
		raddr net.Addr
	}

	tests := []struct {
		name string
		args args
		want *FakeUDPConn
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

			got := NewFakeUDPConn(tt.args.l, tt.args.laddr, tt.args.raddr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFakeUDPConn() = %v, want %v", got, tt.want)
			}
		})
	}
}
