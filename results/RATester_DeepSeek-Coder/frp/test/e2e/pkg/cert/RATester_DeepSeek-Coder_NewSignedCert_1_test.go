package cert

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"

	"k8s.io/client-go/util/cert"
)

func TestNewSignedCert_1(t *testing.T) {
	type args struct {
		cfg    cert.Config
		key    crypto.Signer
		caCert *x509.Certificate
		caKey  crypto.Signer
	}
	tests := []struct {
		name    string
		args    args
		want    *x509.Certificate
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

			got, err := NewSignedCert(tt.args.cfg, tt.args.key, tt.args.caCert, tt.args.caKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSignedCert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSignedCert() = %v, want %v", got, tt.want)
			}
		})
	}
}
