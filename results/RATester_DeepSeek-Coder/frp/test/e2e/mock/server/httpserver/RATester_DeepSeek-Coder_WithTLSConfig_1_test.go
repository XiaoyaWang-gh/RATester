package httpserver

import (
	"crypto/tls"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestWithTLSConfig_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		option   Option
		expected *tls.Config
	}{
		{
			name: "With valid TLS config",
			option: WithTLSConfig(&tls.Config{
				MinVersion: tls.VersionTLS12,
			}),
			expected: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
		{
			name:     "With nil TLS config",
			option:   WithTLSConfig(nil),
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

			s := &Server{}
			s = tc.option(s)
			assert.Equal(t, tc.expected, s.tlsConfig)
		})
	}
}
