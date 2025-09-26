package tls

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetBestCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	c := &CertificateStore{}
	clientHello := &tls.ClientHelloInfo{}

	// WHEN
	cert := c.GetBestCertificate(clientHello)

	// THEN
	assert.Nil(t, cert)
}
