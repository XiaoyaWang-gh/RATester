package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetTruncatedCertificateName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Certificate{}

	c.CertFile = types.FileOrContent("certificateHeader")

	assert.Equal(t, "certificateHeader", c.GetTruncatedCertificateName())
}
