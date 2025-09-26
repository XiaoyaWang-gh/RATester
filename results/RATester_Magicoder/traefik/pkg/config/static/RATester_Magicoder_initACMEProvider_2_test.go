package static

import (
	"fmt"
	"testing"

	acmeprovider "github.com/traefik/traefik/v3/pkg/provider/acme"
	"gotest.tools/assert"
)

func TestinitACMEProvider_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &Configuration{
		CertificatesResolvers: map[string]CertificateResolver{
			"resolver1": {
				ACME: &acmeprovider.Configuration{
					CAServer: "http://example.com/acme",
				},
			},
			"resolver2": {
				ACME: &acmeprovider.Configuration{
					CAServer: "http://example.com/acme",
				},
			},
		},
	}

	config.initACMEProvider()

	for _, resolver := range config.CertificatesResolvers {
		if resolver.ACME != nil {
			assert.Equal(t, "https://example.com/acme", resolver.ACME.CAServer)
		}
	}
}
