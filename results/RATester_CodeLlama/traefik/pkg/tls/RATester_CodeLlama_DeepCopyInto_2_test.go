package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestDeepCopyInto_2(t *testing.T) {
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
	out := &Store{}
	in.DeepCopyInto(out)
	assert.Equal(t, in, out)
}
