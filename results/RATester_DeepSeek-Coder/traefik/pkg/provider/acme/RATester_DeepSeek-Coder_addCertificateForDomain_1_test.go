package acme

import (
	"fmt"
	"testing"

	"github.com/go-acme/lego/v4/certificate"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestAddCertificateForDomain_1(t *testing.T) {
	type args struct {
		domain   types.Domain
		crt      *certificate.Resource
		tlsStore string
	}
	tests := []struct {
		name    string
		p       *Provider
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

			if err := tt.p.addCertificateForDomain(tt.args.domain, tt.args.crt, tt.args.tlsStore); (err != nil) != tt.wantErr {
				t.Errorf("Provider.addCertificateForDomain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
