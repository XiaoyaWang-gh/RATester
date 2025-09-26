package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestRemoteAddr_2(t *testing.T) {
	testCases := []struct {
		name       string
		remoteAddr net.Addr
		expected   net.Addr
	}{
		{
			name:       "Test with valid remote address",
			remoteAddr: &net.IPAddr{IP: net.IPv4(192, 0, 2, 1), Zone: ""},
			expected:   &net.IPAddr{IP: net.IPv4(192, 0, 2, 1), Zone: ""},
		},
		{
			name:       "Test with nil remote address",
			remoteAddr: nil,
			expected:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			conn := &FakeUDPConn{remoteAddr: tc.remoteAddr}
			result := conn.RemoteAddr()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
