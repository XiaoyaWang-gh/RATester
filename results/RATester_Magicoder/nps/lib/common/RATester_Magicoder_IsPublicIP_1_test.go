package common

import (
	"fmt"
	"net"
	"testing"
)

func TestIsPublicIP_1(t *testing.T) {
	tests := []struct {
		name string
		IP   net.IP
		want bool
	}{
		{
			name: "Public IP",
			IP:   net.ParseIP("8.8.8.8"),
			want: true,
		},
		{
			name: "Private IP",
			IP:   net.ParseIP("192.168.1.1"),
			want: false,
		},
		{
			name: "Loopback IP",
			IP:   net.ParseIP("127.0.0.1"),
			want: false,
		},
		{
			name: "Link Local Multicast IP",
			IP:   net.ParseIP("ff02::1"),
			want: false,
		},
		{
			name: "Link Local Unicast IP",
			IP:   net.ParseIP("fe80::1"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsPublicIP(tt.IP); got != tt.want {
				t.Errorf("IsPublicIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
