package ecs

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetIPPort_1(t *testing.T) {
	testCases := []struct {
		name         string
		instance     ecsInstance
		serverPort   string
		expectedIP   string
		expectedPort string
		expectedErr  error
	}{
		{
			name: "success",
			instance: ecsInstance{
				machine: &machine{
					privateIP: "192.0.2.0",
				},
			},
			serverPort:   "8080",
			expectedIP:   "192.0.2.0",
			expectedPort: "8080",
			expectedErr:  nil,
		},
		{
			name: "failure",
			instance: ecsInstance{
				machine: &machine{
					privateIP: "",
				},
			},
			serverPort:   "8080",
			expectedIP:   "",
			expectedPort: "8080",
			expectedErr:  fmt.Errorf("unable to find the IP address for the instance %q: the server is ignored", ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &Provider{}
			ip, port, err := provider.getIPPort(tc.instance, tc.serverPort)
			assert.Equal(t, tc.expectedIP, ip)
			assert.Equal(t, tc.expectedPort, port)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
