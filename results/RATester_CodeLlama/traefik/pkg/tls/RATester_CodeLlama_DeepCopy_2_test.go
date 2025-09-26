package tls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestDeepCopy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := &Store{
		DefaultCertificate: &Certificate{
			CertFile: "certFile",
			KeyFile:  "keyFile",
		},
		DefaultGeneratedCert: &GeneratedCert{
			Resolver: "resolver",
			Domain: &types.Domain{
				Main: "main",
			},
		},
	}
	out := in.DeepCopy()
	assert.Equal(t, in, out)
	assert.NotSame(t, in, out)
	assert.NotSame(t, in.DefaultCertificate, out.DefaultCertificate)
	assert.NotSame(t, in.DefaultGeneratedCert, out.DefaultGeneratedCert)
	assert.NotSame(t, in.DefaultGeneratedCert.Domain, out.DefaultGeneratedCert.Domain)
}
