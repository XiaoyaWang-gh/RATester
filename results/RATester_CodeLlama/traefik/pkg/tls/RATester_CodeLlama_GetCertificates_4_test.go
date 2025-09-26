package tls

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetCertificates_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var c Certificates

	// when
	certs := c.GetCertificates()

	// then
	assert.Equal(t, []tls.Certificate{}, certs)
}
