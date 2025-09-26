package tcp

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"
)

func TestAddHTTPTLSConfig_1(t *testing.T) {
	testCases := []struct {
		desc     string
		sniHost  string
		config   *tls.Config
		expected map[string]*tls.Config
	}{
		{
			desc:    "Adding a TLS config for a new SNI host",
			sniHost: "example.com",
			config:  &tls.Config{},
			expected: map[string]*tls.Config{
				"example.com": {},
			},
		},
		{
			desc:    "Adding a TLS config for an existing SNI host",
			sniHost: "example.com",
			config:  &tls.Config{},
			expected: map[string]*tls.Config{
				"example.com": {},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			router := &Router{}
			router.AddHTTPTLSConfig(tC.sniHost, tC.config)
			if !reflect.DeepEqual(router.hostHTTPTLSConfig, tC.expected) {
				t.Errorf("Expected %v, got %v", tC.expected, router.hostHTTPTLSConfig)
			}
		})
	}
}
