package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestSetRemoteAddr_1(t *testing.T) {
	testCases := []struct {
		name     string
		addr     net.Addr
		expected net.Addr
	}{
		{
			name:     "Test with IP address",
			addr:     &net.IPAddr{IP: net.ParseIP("192.0.2.1")},
			expected: &net.IPAddr{IP: net.ParseIP("192.0.2.1")},
		},
		{
			name:     "Test with Unix address",
			addr:     &net.UnixAddr{Name: "unix", Net: "unix"},
			expected: &net.UnixAddr{Name: "unix", Net: "unix"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			conn := &WrapReadWriteCloserConn{}
			conn.SetRemoteAddr(tc.addr)
			if !reflect.DeepEqual(conn.remoteAddr, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, conn.remoteAddr)
			}
		})
	}
}
