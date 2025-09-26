package group

import (
	"fmt"
	"sync"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestRegister_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	g := &HTTPGroup{
		group:           "",
		groupKey:        "",
		domain:          "",
		location:        "",
		routeByHTTPUser: "",
		createFuncs:     make(map[string]vhost.CreateConnFunc),
		pxyNames:        []string{},
		index:           0,
		ctl:             &HTTPGroupController{},
		mu:              sync.RWMutex{},
	}

	routeConfig := vhost.RouteConfig{
		Domain:                 "test.com",
		Location:               "/",
		RewriteHost:            "",
		Username:               "",
		Password:               "",
		Headers:                make(map[string]string),
		ResponseHeaders:        make(map[string]string),
		RouteByHTTPUser:        "",
		CreateConnFn:           nil,
		ChooseEndpointFn:       nil,
		CreateConnByEndpointFn: nil,
	}

	err := g.Register("proxy1", "group1", "key1", routeConfig)
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}

	err = g.Register("proxy2", "group1", "key1", routeConfig)
	if err == nil {
		t.Errorf("Register should have failed with ErrGroupParamsInvalid")
	}

	err = g.Register("proxy1", "group1", "key2", routeConfig)
	if err == nil {
		t.Errorf("Register should have failed with ErrGroupAuthFailed")
	}

	err = g.Register("proxy3", "group1", "key1", routeConfig)
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}
}
