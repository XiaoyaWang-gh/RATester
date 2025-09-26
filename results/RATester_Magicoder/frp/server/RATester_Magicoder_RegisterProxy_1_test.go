package server

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/server/controller"
	"github.com/fatedier/frp/server/proxy"
)

func TestRegisterProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{
		rc:         &controller.ResourceController{},
		pxyManager: &proxy.Manager{},
		serverCfg:  &v1.ServerConfig{},
	}

	pxyMsg := &msg.NewProxy{
		ProxyName: "testProxy",
		ProxyType: "tcp",
	}

	remoteAddr, err := ctl.RegisterProxy(pxyMsg)
	if err != nil {
		t.Errorf("RegisterProxy failed: %v", err)
	}

	if remoteAddr == "" {
		t.Errorf("RegisterProxy failed: remoteAddr is empty")
	}

	if !ctl.pxyManager.Exist(pxyMsg.ProxyName) {
		t.Errorf("RegisterProxy failed: proxy not added to manager")
	}
}
