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
			name: "Private IP Range 1",
			ip:   net.IP{10, 0, 0, 1},
			want: true,
		},
		{
			name: "Private IP Range 2",
			ip:   net.IP{172, 16, 0, 1},
			want: true,
		},
		{
			name: "Private IP Range 3",
			ip:   net.IP{192, 168, 0, 1},
			want: true,
		},
		{
			name: "Private IP Range 4",
			ip:   net.IP{192, 169, 0, 1},
			want: false,
		},
		{
			name: "IPv6 Private IP Range",
			ip:   net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			want: true,
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
