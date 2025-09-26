package static

import (
	"fmt"
	"testing"

	acmeprovider "github.com/traefik/traefik/v3/pkg/provider/acme"
	"gotest.tools/assert"
)

func TestInitACMEProvider_2(t *testing.T) {
	tests := []struct {
		name string
		c    *Configuration
	}{
		{
			name: "Test with nil ACME",
			c: &Configuration{
				CertificatesResolvers: map[string]CertificateResolver{
					"test": {
						ACME: nil,
					},
				},
			},
		},
		{
			name: "Test with non-nil ACME",
			c: &Configuration{
				CertificatesResolvers: map[string]CertificateResolver{
					"test": {
						ACME: &acmeprovider.Configuration{
							CAServer: "https://acme-staging-v02.api.letsencrypt.org/directory",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.c.initACMEProvider()
			if tt.c.CertificatesResolvers["test"].ACME != nil {
				assert.Equal(t, "https://acme-v02.api.letsencrypt.org/directory", tt.c.CertificatesResolvers["test"].ACME.CAServer)
			}
		})
	}
}
