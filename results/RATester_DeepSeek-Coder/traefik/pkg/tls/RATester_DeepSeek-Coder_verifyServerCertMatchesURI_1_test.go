package tls

import (
	"crypto/x509"
	"fmt"
	"testing"
)

func TestVerifyServerCertMatchesURI_1(t *testing.T) {
	type args struct {
		uri  string
		cert *x509.Certificate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := verifyServerCertMatchesURI(tt.args.uri, tt.args.cert)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyServerCertMatchesURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
