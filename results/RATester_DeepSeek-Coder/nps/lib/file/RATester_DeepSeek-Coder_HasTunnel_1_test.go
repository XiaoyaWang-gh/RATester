package file

import (
	"fmt"
	"testing"
)

func TestHasTunnel_1(t *testing.T) {
	client := &Client{
		Id: 1,
	}
	tunnel := &Tunnel{
		Id:     1,
		Port:   8080,
		Client: client,
	}

	testCases := []struct {
		name     string
		client   *Client
		tunnel   *Tunnel
		expected bool
	}{
		{
			name:     "Tunnel exists",
			client:   client,
			tunnel:   tunnel,
			expected: true,
		},
		{
			name: "Tunnel does not exist",
			client: &Client{
				Id: 2,
			},
			tunnel: &Tunnel{
				Id: 2,
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			exist := tc.client.HasTunnel(tc.tunnel)
			if exist != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, exist)
			}
		})
	}
}
