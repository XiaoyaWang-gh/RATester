package httplib

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetTLSClientConfig_1(t *testing.T) {
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

			b := &BeegoHTTPRequest{}
			b.SetTLSClientConfig(tc.config)

			if b.setting.TLSClientConfig != tc.config {
				t.Errorf("Expected TLSClientConfig to be %v, but got %v", tc.config, b.setting.TLSClientConfig)
			}
		})
	}
}
