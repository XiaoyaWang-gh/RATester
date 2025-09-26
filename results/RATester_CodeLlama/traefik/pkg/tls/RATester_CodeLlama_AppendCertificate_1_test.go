package tls

import (
	"crypto/tls"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppendCertificate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Certificate{
		CertFile: "certFile",
		KeyFile:  "keyFile",
	}
	certs := make(map[string]map[string]*tls.Certificate)
	storeName := "storeName"

	// when
	err := c.AppendCertificate(certs, storeName)

	// then
	require.NoError(t, err)
	require.NotNil(t, certs[storeName])
	require.Equal(t, 1, len(certs[storeName]))
	require.NotNil(t, certs[storeName]["certFile"])
}
