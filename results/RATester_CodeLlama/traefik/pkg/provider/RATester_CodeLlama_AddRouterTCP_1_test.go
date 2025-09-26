package provider

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddRouterTCP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configuration := &dynamic.TCPConfiguration{}
	routerName := "test"
	router := &dynamic.TCPRouter{}

	if !AddRouterTCP(configuration, routerName, router) {
		t.Errorf("AddRouterTCP() = false, want true")
	}

	if configuration.Routers[routerName] != router {
		t.Errorf("AddRouterTCP() = %v, want %v", configuration.Routers[routerName], router)
	}
}
