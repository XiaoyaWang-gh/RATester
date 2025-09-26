package v1

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestUnmarshalFromMsg_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    *msg.NewProxy
		expected SUDPProxyConfig
	}{
		{
			name: "Test case 1",
			input: &msg.NewProxy{
				ProxyName: "testProxy",
				ProxyType: "SUDP",
				Sk:        "testSecretKey",
				AllowUsers: []string{
					"user1",
					"user2",
				},
			},
			expected: SUDPProxyConfig{
				ProxyBaseConfig: ProxyBaseConfig{
					Name: "testProxy",
					Type: "SUDP",
				},
				Secretkey:  "testSecretKey",
				AllowUsers: []string{"user1", "user2"},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			config := &SUDPProxyConfig{}
			config.UnmarshalFromMsg(tc.input)

			if config.Name != tc.expected.Name {
				t.Errorf("Expected Name: %s, but got: %s", tc.expected.Name, config.Name)
			}

			if config.Type != tc.expected.Type {
				t.Errorf("Expected Type: %s, but got: %s", tc.expected.Type, config.Type)
			}

			if config.Secretkey != tc.expected.Secretkey {
				t.Errorf("Expected Secretkey: %s, but got: %s", tc.expected.Secretkey, config.Secretkey)
			}

			if !reflect.DeepEqual(config.AllowUsers, tc.expected.AllowUsers) {
				t.Errorf("Expected AllowUsers: %v, but got: %v", tc.expected.AllowUsers, config.AllowUsers)
			}
		})
	}
}
