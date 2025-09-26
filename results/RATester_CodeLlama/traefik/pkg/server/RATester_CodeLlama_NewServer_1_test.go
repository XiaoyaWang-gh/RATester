package server

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func TestNewServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	routinesPool := &safe.Pool{}
	entryPoints := TCPEntryPoints{}
	entryPointsUDP := UDPEntryPoints{}
	watcher := &ConfigurationWatcher{}
	observabilityMgr := &middleware.ObservabilityMgr{}

	// when
	server := NewServer(routinesPool, entryPoints, entryPointsUDP, watcher, observabilityMgr)

	// then
	require.NotNil(t, server)
	require.Equal(t, routinesPool, server.routinesPool)
	require.Equal(t, entryPoints, server.tcpEntryPoints)
	require.Equal(t, entryPointsUDP, server.udpEntryPoints)
	require.Equal(t, watcher, server.watcher)
	require.Equal(t, observabilityMgr, server.observabilityMgr)
}
