package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddRouterUDP_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configuration := &dynamic.UDPConfiguration{}
	routerName := "test"
	router := &dynamic.UDPRouter{
		EntryPoints: []string{"test"},
		Service:     "test",
	}

	if !AddRouterUDP(configuration, routerName, router) {
		t.Errorf("AddRouterUDP() = false, want true")
	}

	if configuration.Routers[routerName] == nil {
		t.Errorf("configuration.Routers[routerName] = nil, want &dynamic.UDPRouter")
	}

	if !reflect.DeepEqual(configuration.Routers[routerName], router) {
		t.Errorf("configuration.Routers[routerName] = %+v, want %+v", configuration.Routers[routerName], router)
	}
}
