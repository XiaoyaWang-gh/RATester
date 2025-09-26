package tls

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCertificateFromBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Certificate{
		CertFile: "certFile",
		KeyFile:  "keyFile",
	}

	cert, err := c.GetCertificateFromBytes()
	require.NoError(t, err)
	require.NotNil(t, cert)
}
