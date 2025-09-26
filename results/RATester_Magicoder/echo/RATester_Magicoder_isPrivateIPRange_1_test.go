package echo

import (
	"fmt"
	"net"
	"testing"
)

func TestIsPrivateIPRange_1(t *testing.T) {
	tests := []struct {
		name string
		ip   net.IP
		want bool
	}{
		{
			name: "Private IP Range",
			ip:   net.ParseIP("10.0.0.1"),
			want: true,
		},
		{
			name: "Private IP Range",
			ip:   net.ParseIP("172.16.0.1"),
			want: true,
		},
		{
			name: "Private IP Range",
			ip:   net.ParseIP("192.168.0.1"),
			want: true,
		},
		{
			name: "Public IP Range",
			ip:   net.ParseIP("8.8.8.8"),
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

			if got := isPrivateIPRange(tt.ip); got != tt.want {
				t.Errorf("isPrivateIPRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
