package ingress

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/tls"
	netv1 "k8s.io/api/networking/v1"
)

func TestGetCertificates_3(t *testing.T) {
	type args struct {
		ctx        context.Context
		ingress    *netv1.Ingress
		k8sClient  Client
		tlsConfigs map[string]*tls.CertAndStores
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

			if err := getCertificates(tt.args.ctx, tt.args.ingress, tt.args.k8sClient, tt.args.tlsConfigs); (err != nil) != tt.wantErr {
				t.Errorf("getCertificates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
