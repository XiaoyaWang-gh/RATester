package tailscale

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestFetchCerts_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	p.certByDomain = make(map[string]traefiktls.Certificate)

	ctx := context.Background()
	domains := []string{"example.com", "example.net"}

	p.fetchCerts(ctx, domains)

	assert.Equal(t, 2, len(p.certByDomain))
	assert.Equal(t, traefiktls.Certificate{
		CertFile: types.FileOrContent("cert"),
		KeyFile:  types.FileOrContent("key"),
	}, p.certByDomain["example.com"])
	assert.Equal(t, traefiktls.Certificate{
		CertFile: types.FileOrContent("cert"),
		KeyFile:  types.FileOrContent("key"),
	}, p.certByDomain["example.net"])
}
