package static

import (
	"fmt"
	"testing"

	acmeprovider "github.com/traefik/traefik/v3/pkg/provider/acme"
)

func TestValidateConfiguration_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *Configuration
		wantErr bool
	}{
		{
			name: "valid configuration",
			config: &Configuration{
				CertificatesResolvers: map[string]CertificateResolver{
					"resolver1": {
						ACME: &acmeprovider.Configuration{
							Storage: "path/to/storage",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid configuration with ACME and Tailscale",
			config: &Configuration{
				CertificatesResolvers: map[string]CertificateResolver{
					"resolver1": {
						ACME: &acmeprovider.Configuration{
							Storage: "path/to/storage",
						},
						Tailscale: &struct{}{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid configuration with no storage location for ACME",
			config: &Configuration{
				CertificatesResolvers: map[string]CertificateResolver{
					"resolver1": {
						ACME: &acmeprovider.Configuration{},
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tt.config.ValidateConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
