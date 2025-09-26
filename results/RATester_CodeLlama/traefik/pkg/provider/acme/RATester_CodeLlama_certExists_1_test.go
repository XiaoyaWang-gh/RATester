package acme

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestCertExists_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.certificates = []*CertAndStore{
		{
			Certificate: Certificate{
				Domain: types.Domain{
					Main: "test.com",
				},
			},
		},
	}
	validDomains := []string{"test.com"}
	assert.True(t, p.certExists(validDomains))
}
