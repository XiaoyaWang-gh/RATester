package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestnewUdpConnByAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    *net.UDPConn
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				addr: "127.0.0.1:8080",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				addr: "invalid_address",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newUdpConnByAddr(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("newUdpConnByAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUdpConnByAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
