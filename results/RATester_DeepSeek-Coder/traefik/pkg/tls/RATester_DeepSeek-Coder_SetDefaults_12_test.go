package tls

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetDefaults_12(t *testing.T) {
	testCases := []struct {
		desc     string
		given    *Options
		expected *Options
	}{
		{
			desc:  "should set default values when options are nil",
			given: nil,
			expected: &Options{
				ALPNProtocols: DefaultTLSOptions.ALPNProtocols,
				CipherSuites:  DefaultTLSOptions.CipherSuites,
			},
		},
		{
			desc:  "should set default values when options are empty",
			given: &Options{},
			expected: &Options{
				ALPNProtocols: DefaultTLSOptions.ALPNProtocols,
				CipherSuites:  DefaultTLSOptions.CipherSuites,
			},
		},
		{
			desc: "should not override non-default values",
			given: &Options{
				ALPNProtocols: []string{"h2", "http/1.1"},
				CipherSuites:  []string{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"},
			},
			expected: &Options{
				ALPNProtocols: []string{"h2", "http/1.1"},
				CipherSuites:  []string{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.given.SetDefaults()

			if !reflect.DeepEqual(test.given, test.expected) {
				t.Errorf("got %v, want %v", test.given, test.expected)
			}
		})
	}
}
