package request

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestTLSConfig_1(t *testing.T) {
	testCases := []struct {
		name      string
		config    *tls.Config
		expectErr bool
	}{
		{
			name: "valid config",
			config: &tls.Config{
				InsecureSkipVerify: true,
			},
			expectErr: false,
		},
		{
			name:      "nil config",
			config:    nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req := &Request{}
			err := req.TLSConfig(tc.config)
			if tc.expectErr && err == nil {
				t.Errorf("expected error, but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}
