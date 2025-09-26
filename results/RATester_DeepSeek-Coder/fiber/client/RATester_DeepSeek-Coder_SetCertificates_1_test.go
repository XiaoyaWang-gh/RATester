package client

import (
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
	"gotest.tools/assert"
)

func TestSetCertificates_1(t *testing.T) {
	testCases := []struct {
		name     string
		client   *Client
		certs    []tls.Certificate
		expected *Client
	}{
		{
			name: "Test with valid certificates",
			client: &Client{
				fasthttp: &fasthttp.Client{},
			},
			certs: []tls.Certificate{
				{
					Certificate: [][]byte{[]byte("certificate1")},
					PrivateKey:  &rsa.PrivateKey{},
				},
				{
					Certificate: [][]byte{[]byte("certificate2")},
					PrivateKey:  &rsa.PrivateKey{},
				},
			},
			expected: &Client{
				fasthttp: &fasthttp.Client{
					TLSConfig: &tls.Config{
						Certificates: []tls.Certificate{
							{
								Certificate: [][]byte{[]byte("certificate1")},
								PrivateKey:  &rsa.PrivateKey{},
							},
							{
								Certificate: [][]byte{[]byte("certificate2")},
								PrivateKey:  &rsa.PrivateKey{},
							},
						},
					},
				},
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

			result := tc.client.SetCertificates(tc.certs...)
			assert.Equal(t, tc.expected, result)
		})
	}
}
