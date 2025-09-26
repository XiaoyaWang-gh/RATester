package client

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetCertificates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given: A Client
	c := &Client{}

	// and: A list of certificates
	certs := []tls.Certificate{}

	// when: SetCertificates is called
	c.SetCertificates(certs...)

	// then: The certificates are set
	assert.Equal(t, certs, c.TLSConfig().Certificates)
}
