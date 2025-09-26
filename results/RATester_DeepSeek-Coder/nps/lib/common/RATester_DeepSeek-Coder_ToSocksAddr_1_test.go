package common

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestToSocksAddr_1(t *testing.T) {
	type args struct {
		addr net.Addr
	}
	tests := []struct {
		name string
		args args
		want *Addr
	}{
		{
			name: "Test with nil addr",
			args: args{
				addr: nil,
			},
			want: &Addr{
				Type: ipV4,
				Host: "0.0.0.0",
				Port: 0,
			},
		},
		{
			name: "Test with valid addr",
			args: args{
				addr: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 22,
					Zone: "",
				},
			},
			want: &Addr{
				Type: ipV4,
				Host: "192.0.2.1",
				Port: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ToSocksAddr(tt.args.addr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSocksAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
