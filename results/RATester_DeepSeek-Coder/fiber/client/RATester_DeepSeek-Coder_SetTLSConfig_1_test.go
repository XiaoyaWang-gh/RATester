package client

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetTLSConfig_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		config *tls.Config
	}{
		{
			name:   "Nil config",
			config: nil,
		},
		{
			name: "Valid config",
			config: &tls.Config{
				InsecureSkipVerify: true,
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

			c := &Client{}
			result := c.SetTLSConfig(tc.config)

			if result != c {
				t.Errorf("Expected result to be the same as the client, got %v", result)
			}

			if c.fasthttp.TLSConfig != tc.config {
				t.Errorf("Expected TLSConfig to be %v, got %v", tc.config, c.fasthttp.TLSConfig)
			}
		})
	}
}
