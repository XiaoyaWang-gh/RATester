package server

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/server/proxy"
)

func TestRegisterProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{
		serverCfg: &v1.ServerConfig{
			MaxPortsPerClient: 10,
		},
		proxies: make(map[string]proxy.Proxy),
	}

	pxyMsg := &msg.NewProxy{
		ProxyName: "test",
		ProxyType: "tcp",
	}

	_, err := ctl.RegisterProxy(pxyMsg)
	if err != nil {
		t.Errorf("RegisterProxy failed: %v", err)
	}

	if _, exists := ctl.proxies["test"]; !exists {
		t.Errorf("Proxy not added to proxies map")
	}
}
