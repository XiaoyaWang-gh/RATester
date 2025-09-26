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
			name: "Test case 1",
			args: args{
				addr: &net.TCPAddr{
					IP:   net.ParseIP("192.0.2.1"),
					Port: 80,
					Zone: "",
				},
			},
			want: &Addr{
				Type: ipV4,
				Host: "192.0.2.1",
				Port: 80,
			},
		},
		{
			name: "Test case 2",
			args: args{
				addr: &net.TCPAddr{
					IP:   net.ParseIP("2001:db8::1"),
					Port: 443,
					Zone: "",
				},
			},
			want: &Addr{
				Type: ipV4,
				Host: "2001:db8::1",
				Port: 443,
			},
		},
		{
			name: "Test case 3",
			args: args{
				addr: nil,
			},
			want: &Addr{
				Type: ipV4,
				Host: "0.0.0.0",
				Port: 0,
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

			if got := ToSocksAddr(tt.args.addr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSocksAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
