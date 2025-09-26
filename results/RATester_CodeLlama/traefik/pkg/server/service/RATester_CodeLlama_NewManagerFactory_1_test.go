package service

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/server/middleware"
)

func TestNewManagerFactory_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	staticConfiguration := static.Configuration{}
	routinesPool := &safe.Pool{}
	observabilityMgr := &middleware.ObservabilityMgr{}
	roundTripperManager := &RoundTripperManager{}
	acmeHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// when
	managerFactory := NewManagerFactory(staticConfiguration, routinesPool, observabilityMgr, roundTripperManager, acmeHTTPHandler)

	// then
	require.NotNil(t, managerFactory)
}
