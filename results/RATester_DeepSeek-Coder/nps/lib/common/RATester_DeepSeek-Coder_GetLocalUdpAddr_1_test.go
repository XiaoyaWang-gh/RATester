package common

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetLocalUdpAddr_1(t *testing.T) {
	type args struct {
		conn net.Conn
		err  error
	}
	tests := []struct {
		name    string
		args    args
		want    net.Conn
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test case 1",
			args: args{
				conn: &net.UDPConn{},
				err:  nil,
			},
			want:    &net.UDPConn{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := GetLocalUdpAddr()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalUdpAddr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocalUdpAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
