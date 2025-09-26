package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewKCPConnFromUDP_1(t *testing.T) {
	type args struct {
		conn      *net.UDPConn
		connected bool
		raddr     string
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
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

			got, err := NewKCPConnFromUDP(tt.args.conn, tt.args.connected, tt.args.raddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKCPConnFromUDP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKCPConnFromUDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
