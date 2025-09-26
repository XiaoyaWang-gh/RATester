package client

import (
	"fmt"
	"testing"
)

func TestNewUdpConnByAddr_1(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid address",
			args: args{
				addr: "127.0.0.1:8080",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid address",
			args: args{
				addr: "127.0.0.1:8081",
			},
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

			_, err := newUdpConnByAddr(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("newUdpConnByAddr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
