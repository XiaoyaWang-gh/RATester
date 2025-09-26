package client

import (
	"fmt"
	"net"
	"testing"
)

func TestGetRemoteAddressFromServer_1(t *testing.T) {
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
			"case 1",
			args{
				rAddr:       "127.0.0.1:8080",
				localConn:   &net.UDPConn{},
				md5Password: "123456",
				role:        "1",
				add:         1,
			},
			false,
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
