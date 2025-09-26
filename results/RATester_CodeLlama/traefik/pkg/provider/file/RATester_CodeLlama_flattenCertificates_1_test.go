package file

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestFlattenCertificates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	tlsConfig := &dynamic.TLSConfiguration{}
	certs := flattenCertificates(ctx, tlsConfig)
	assert.Equal(t, 0, len(certs))
}
