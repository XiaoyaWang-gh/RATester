package cert

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"reflect"
	"testing"
)

func TestValidCACert_1(t *testing.T) {
	type args struct {
		caKey  []byte
		caCert []byte
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 *rsa.PrivateKey
		want2 *x509.Certificate
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

			cp := &SelfSignedCertGenerator{
				caKey:  tt.args.caKey,
				caCert: tt.args.caCert,
			}
			got, got1, got2 := cp.validCACert()
			if got != tt.want {
				t.Errorf("validCACert() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validCACert() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("validCACert() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
