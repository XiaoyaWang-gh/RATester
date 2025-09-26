package passtlsclientcert

import (
	"context"
	"crypto/x509"
	"fmt"
	"testing"
)

func TestGetCertificates_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		certs []*x509.Certificate
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getCertificates(tt.args.ctx, tt.args.certs); got != tt.want {
				t.Errorf("getCertificates() = %v, want %v", got, tt.want)
			}
		})
	}
}
