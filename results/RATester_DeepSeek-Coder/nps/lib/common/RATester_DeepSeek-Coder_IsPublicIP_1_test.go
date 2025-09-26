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
			name: "Public IPv4",
			IP:   net.ParseIP("8.8.8.8"),
			want: true,
		},
		{
			name: "Private IPv4",
			IP:   net.ParseIP("192.168.1.1"),
			want: false,
		},
		{
			name: "Loopback IPv4",
			IP:   net.ParseIP("127.0.0.1"),
			want: false,
		},
		{
			name: "Link Local IPv4",
			IP:   net.ParseIP("169.254.169.254"),
			want: false,
		},
		{
			name: "IPv6",
			IP:   net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"),
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
