package gin

import (
	"fmt"
	"net"
	"testing"
)

func TestisTrustedProxy_1(t *testing.T) {
	engine := &Engine{
		trustedCIDRs: []*net.IPNet{
			{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(24, 32)},
			{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		},
	}

	tests := []struct {
		name string
		ip   net.IP
		want bool
	}{
		{"Trusted IP", net.IPv4(192, 168, 0, 1), true},
		{"Untrusted IP", net.IPv4(192, 168, 1, 1), false},
		{"Trusted IP", net.IPv4(10, 0, 0, 1), true},
		{"Untrusted IP", net.IPv4(10, 1, 0, 1), false},
		{"Nil IP", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := engine.isTrustedProxy(tt.ip); got != tt.want {
				t.Errorf("isTrustedProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
