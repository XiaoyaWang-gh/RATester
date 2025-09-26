package gin

import (
	"fmt"
	"net"
	"testing"
)

func TestIsTrustedProxy_1(t *testing.T) {
	engine := &Engine{
		trustedCIDRs: []*net.IPNet{
			{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(24, 32)},
			{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		},
	}

	testCases := []struct {
		name     string
		ip       net.IP
		expected bool
	}{
		{"Trusted IP", net.IPv4(192, 168, 0, 1), true},
		{"Not trusted IP", net.IPv4(192, 168, 1, 1), false},
		{"Trusted IP with different mask", net.IPv4(10, 0, 0, 1), true},
		{"Not trusted IP with different mask", net.IPv4(10, 1, 0, 0), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := engine.isTrustedProxy(tc.ip)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
