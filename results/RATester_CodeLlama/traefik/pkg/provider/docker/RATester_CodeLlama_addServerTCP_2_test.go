package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddServerTCP_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	p := &DynConfBuilder{}
	container := dockerData{}
	loadBalancer := &dynamic.TCPServersLoadBalancer{}
	err := p.addServerTCP(ctx, container, loadBalancer)
	require.Error(t, err)
}
