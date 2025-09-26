package vhost

import (
	"fmt"
	"testing"
)

func TestUnRegister_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rp := &HTTPReverseProxy{
		vhostRouter: &Routers{
			indexByDomain: make(map[string]routerByHTTPUser),
		},
	}

	routeCfg := RouteConfig{
		Domain:   "example.com",
		Location: "/api",
	}

	rp.vhostRouter.Add(routeCfg.Domain, routeCfg.Location, "", struct{}{})

	rp.UnRegister(routeCfg)

	_, exist := rp.vhostRouter.Get(routeCfg.Domain, routeCfg.Location, "")
	if exist {
		t.Errorf("Expected route to be unregistered")
	}
}
