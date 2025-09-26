package tls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"testing"
)

func TestVerifyPeerCertificate_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		uri           string
		cfg           *tls.Config
		rawCerts      [][]byte
		expectedError error
	}{
		{
			name:          "valid certificate",
			uri:           "example.com",
			cfg:           &tls.Config{RootCAs: x509.NewCertPool()},
			rawCerts:      [][]byte{{1, 2, 3}, {4, 5, 6}},
			expectedError: nil,
		},
		{
			name:          "invalid certificate",
			uri:           "example.com",
			cfg:           &tls.Config{RootCAs: x509.NewCertPool()},
			rawCerts:      [][]byte{{1, 2, 3}, {4, 5, 7}}, // last byte is different
			expectedError: errors.New("some error"),
		},
		{
			name:          "empty uri",
			uri:           "",
			cfg:           &tls.Config{RootCAs: x509.NewCertPool()},
			rawCerts:      [][]byte{{1, 2, 3}, {4, 5, 6}},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := VerifyPeerCertificate(tc.uri, tc.cfg, tc.rawCerts)
			if err != tc.expectedError {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
