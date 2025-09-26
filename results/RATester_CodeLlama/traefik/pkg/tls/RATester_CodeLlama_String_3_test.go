package tls

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestString_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Certificates{}
	assert.Equal(t, "", c.String())

	c = &Certificates{
		{
			CertFile: types.FileOrContent("certFile"),
			KeyFile:  types.FileOrContent("keyFile"),
		},
	}
	assert.Equal(t, "certFile,keyFile", c.String())
}
