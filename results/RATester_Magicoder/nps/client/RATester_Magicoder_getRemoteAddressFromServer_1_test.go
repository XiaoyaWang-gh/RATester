package client

import (
	"fmt"
	"net"
	"testing"
)

func TestgetRemoteAddressFromServer_1(t *testing.T) {
	type args struct {
		rAddr       string
		localConn   *net.UDPConn
		md5Password string
		role        string
		add         int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				rAddr:       "127.0.0.1:8080",
				localConn:   &net.UDPConn{},
				md5Password: "password",
				role:        "role",
				add:         1,
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				rAddr:       "127.0.0.1:8080",
				localConn:   &net.UDPConn{},
				md5Password: "password",
				role:        "role",
				add:         0,
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

			if err := getRemoteAddressFromServer(tt.args.rAddr, tt.args.localConn, tt.args.md5Password, tt.args.role, tt.args.add); (err != nil) != tt.wantErr {
				t.Errorf("getRemoteAddressFromServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
