package net

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestLocalAddr_2(t *testing.T) {
	testCases := []struct {
		name      string
		localAddr net.Addr
		expected  net.Addr
	}{
		{
			name:      "Test case 1",
			localAddr: &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080, Zone: ""},
			expected:  &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8080, Zone: ""},
		},
		{
			name:      "Test case 2",
			localAddr: &net.TCPAddr{IP: net.ParseIP("::1"), Port: 8080, Zone: ""},
			expected:  &net.TCPAddr{IP: net.ParseIP("::1"), Port: 8080, Zone: ""},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fakeConn := &FakeUDPConn{
				localAddr: tc.localAddr,
			}

			result := fakeConn.LocalAddr()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
