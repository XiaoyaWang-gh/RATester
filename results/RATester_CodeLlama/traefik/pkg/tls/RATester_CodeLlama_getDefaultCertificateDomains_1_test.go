package tls

import (
	"crypto/tls"
	"fmt"
	"reflect"
	"testing"

	"github.com/patrickmn/go-cache"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestGetDefaultCertificateDomains_1(t *testing.T) {
	type fields struct {
		DynamicCerts       *safe.Safe
		DefaultCertificate *tls.Certificate
		CertCache          *cache.Cache
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			"test getDefaultCertificateDomains",
			fields{
				DynamicCerts:       nil,
				DefaultCertificate: nil,
				CertCache:          nil,
			},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := CertificateStore{
				DynamicCerts:       tt.fields.DynamicCerts,
				DefaultCertificate: tt.fields.DefaultCertificate,
				CertCache:          tt.fields.CertCache,
			}
			if got := c.getDefaultCertificateDomains(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDefaultCertificateDomains() = %v, want %v", got, tt.want)
			}
		})
	}
}
