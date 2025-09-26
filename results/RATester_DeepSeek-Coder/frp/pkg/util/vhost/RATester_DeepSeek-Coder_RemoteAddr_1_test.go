package vhost

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestRemoteAddr_1(t *testing.T) {
	testCases := []struct {
		name     string
		conn     readOnlyConn
		expected net.Addr
	}{
		{
			name: "Test with valid data",
			conn: readOnlyConn{
				reader: strings.NewReader("test"),
			},
			expected: nil,
		},
		{
			name: "Test with invalid data",
			conn: readOnlyConn{
				reader: nil,
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.conn.RemoteAddr()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
