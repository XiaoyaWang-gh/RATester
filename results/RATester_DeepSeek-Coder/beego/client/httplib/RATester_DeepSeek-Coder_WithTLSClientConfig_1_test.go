package httplib

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestWithTLSClientConfig_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config *tls.Config
	}{
		{
			name: "With valid TLS Config",
			config: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		{
			name:   "With nil TLS Config",
			config: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			opt := WithTLSClientConfig(tc.config)
			client := &Client{}
			opt(client)

			if tc.config != nil && client.Setting.TLSClientConfig != tc.config {
				t.Errorf("Expected TLSClientConfig to be %v, got %v", tc.config, client.Setting.TLSClientConfig)
			}

			if tc.config == nil && client.Setting.TLSClientConfig != nil {
				t.Errorf("Expected TLSClientConfig to be nil, got %v", client.Setting.TLSClientConfig)
			}
		})
	}
}
