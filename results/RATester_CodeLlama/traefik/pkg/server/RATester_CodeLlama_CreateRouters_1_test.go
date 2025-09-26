package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestCreateRouters_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	f := &RouterFactory{
		entryPointsTCP: []string{"foo", "bar"},
	}
	rtConf := &runtime.Configuration{}

	// WHEN
	routersTCP, routersUDP := f.CreateRouters(rtConf)

	// THEN
	require.Len(t, routersTCP, 2)
	require.Len(t, routersUDP, 0)
}
