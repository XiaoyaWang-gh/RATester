package tls

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetDefaults_12(t *testing.T) {
	testCases := []struct {
		desc     string
		input    *Options
		expected *Options
	}{
		{
			desc:  "should set defaults when options are nil",
			input: nil,
			expected: &Options{
				ALPNProtocols: DefaultTLSOptions.ALPNProtocols,
				CipherSuites:  DefaultTLSOptions.CipherSuites,
			},
		},
		{
			desc:  "should set defaults when options are empty",
			input: &Options{},
			expected: &Options{
				ALPNProtocols: DefaultTLSOptions.ALPNProtocols,
				CipherSuites:  DefaultTLSOptions.CipherSuites,
			},
		},
		{
			desc: "should not override non-default values",
			input: &Options{
				ALPNProtocols: []string{"http/1.1", "h2"},
				CipherSuites:  []string{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"},
			},
			expected: &Options{
				ALPNProtocols: []string{"http/1.1", "h2"},
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

			test.input.SetDefaults()

			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("got %v, want %v", test.input, test.expected)
			}
		})
	}
}
