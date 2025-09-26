package tls

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestGetServerCertificates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Manager{
		stores: map[string]*CertificateStore{
			"default": {
				DynamicCerts: &safe.Safe{},
				DefaultCertificate: &tls.Certificate{
					Certificate: [][]byte{[]byte("cert")},
				},
			},
		},
	}

	certificates := m.GetServerCertificates()
	assert.Len(t, certificates, 1)
	assert.Equal(t, []byte("cert"), certificates[0].Raw)
}
