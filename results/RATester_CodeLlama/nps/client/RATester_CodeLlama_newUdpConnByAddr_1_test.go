package client

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestNewUdpConnByAddr_1(t *testing.T) {
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
			name: "test1",
			args: args{
				addr: "127.0.0.1:8080",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test2",
			args: args{
				addr: "127.0.0.1:8081",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test3",
			args: args{
				addr: "127.0.0.1:8082",
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
				t.Errorf("newUdpConnByAddr() got = %v, want %v", got, tt.want)
			}
		})
	}
}
