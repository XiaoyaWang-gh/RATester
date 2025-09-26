package tailscale

import (
	"context"
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestWatchDomains_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	p := &Provider{}
	p.dynConfigs = make(chan dynamic.Configuration)
	p.dynMessages = make(chan dynamic.Message)
	p.certByDomain = make(map[string]traefiktls.Certificate)

	go p.watchDomains(ctx)

	// TODO: add test cases.
	t.Fail()
}
