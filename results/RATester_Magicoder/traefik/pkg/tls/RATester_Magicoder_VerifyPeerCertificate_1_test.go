package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"testing"
)

func TestVerifyPeerCertificate_1(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		cfg      *tls.Config
		rawCerts [][]byte
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Test case 1",
			uri:  "example.com",
			cfg: &tls.Config{
				RootCAs: x509.NewCertPool(),
			},
			rawCerts: [][]byte{[]byte("certificate data")},
			wantErr:  false,
		},
		{
			name: "Test case 2",
			uri:  "example.com",
			cfg: &tls.Config{
				RootCAs: x509.NewCertPool(),
			},
			rawCerts: [][]byte{[]byte("certificate data")},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := VerifyPeerCertificate(tt.uri, tt.cfg, tt.rawCerts); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPeerCertificate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
