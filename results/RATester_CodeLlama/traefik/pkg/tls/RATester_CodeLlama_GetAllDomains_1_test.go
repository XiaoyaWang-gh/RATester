package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestGetAllDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := CertificateStore{
		DynamicCerts: &safe.Safe{},
	}

	// when
	allDomains := c.GetAllDomains()

	// then
	assert.Empty(t, allDomains)
}
