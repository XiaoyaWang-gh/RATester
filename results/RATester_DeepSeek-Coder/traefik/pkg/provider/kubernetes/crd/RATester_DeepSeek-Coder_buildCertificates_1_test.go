package crd

import (
	"fmt"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestBuildCertificates_1(t *testing.T) {
	type args struct {
		client       Client
		tlsStore     string
		namespace    string
		certificates []traefikv1alpha1.Certificate
		tlsConfigs   map[string]*tls.CertAndStores
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

			if err := buildCertificates(tt.args.client, tt.args.tlsStore, tt.args.namespace, tt.args.certificates, tt.args.tlsConfigs); (err != nil) != tt.wantErr {
				t.Errorf("buildCertificates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
