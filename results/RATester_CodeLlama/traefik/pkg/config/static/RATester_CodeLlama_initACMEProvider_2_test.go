package static

import (
	"fmt"
	"testing"

	acmeprovider "github.com/traefik/traefik/v3/pkg/provider/acme"
)

func TestInitACMEProvider_2(t *testing.T) {
	type args struct {
		c *Configuration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				c: &Configuration{
					CertificatesResolvers: map[string]CertificateResolver{
						"test": {
							ACME: &acmeprovider.Configuration{
								CAServer: "test",
							},
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

			tt.args.c.initACMEProvider()
		})
	}
}
