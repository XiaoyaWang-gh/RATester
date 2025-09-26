package crd

import (
	"context"
	"fmt"
	"testing"

	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetTLSTCP_1(t *testing.T) {
	type args struct {
		ctx          context.Context
		ingressRoute *traefikv1alpha1.IngressRouteTCP
		k8sClient    Client
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

			if err := getTLSTCP(tt.args.ctx, tt.args.ingressRoute, tt.args.k8sClient, tt.args.tlsConfigs); (err != nil) != tt.wantErr {
				t.Errorf("getTLSTCP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
