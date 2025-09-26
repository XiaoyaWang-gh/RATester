package echo

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestTrustIPRange_1(t *testing.T) {
	testCases := []struct {
		name     string
		ipRange  *net.IPNet
		expected []*net.IPNet
	}{
		{
			name: "Test with valid IP range",
			ipRange: &net.IPNet{
				IP:   net.IPv4(192, 168, 1, 0),
				Mask: net.CIDRMask(24, 32),
			},
			expected: []*net.IPNet{
				{
					IP:   net.IPv4(192, 168, 1, 0),
					Mask: net.CIDRMask(24, 32),
				},
			},
		},
		{
			name: "Test with invalid IP range",
			ipRange: &net.IPNet{
				IP:   net.IPv4(192, 168, 1, 0),
				Mask: net.CIDRMask(33, 32),
			},
			expected: []*net.IPNet{
				{
					IP:   net.IPv4(192, 168, 1, 0),
					Mask: net.CIDRMask(24, 32),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ipChecker{}
			TrustIPRange(tc.ipRange)(c)
			if !reflect.DeepEqual(c.trustExtraRanges, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, c.trustExtraRanges)
			}
		})
	}
}
