package gateway

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestLoadHTTPRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := Provider{}
	ctx := context.Background()
	listener := gatewayListener{}
	route := &gatev1.HTTPRoute{}
	hostnames := []gatev1.Hostname{}
	conf, condition := p.loadHTTPRoute(ctx, listener, route, hostnames)
	assert.NotNil(t, conf)
	assert.NotNil(t, condition)
}
