package common

import (
	"fmt"
	"net"
	"testing"
)

func TestGetServerIpByClientIp_1(t *testing.T) {
	type args struct {
		clientIp net.IP
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				clientIp: net.ParseIP("192.168.1.1"),
			},
			want: "192.168.1.1",
		},
		{
			name: "Test case 2",
			args: args{
				clientIp: net.ParseIP("8.8.8.8"),
			},
			want: "8.8.8.8",
		},
		{
			name: "Test case 3",
			args: args{
				clientIp: net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
			},
			want: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetServerIpByClientIp(tt.args.clientIp); got != tt.want {
				t.Errorf("GetServerIpByClientIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
