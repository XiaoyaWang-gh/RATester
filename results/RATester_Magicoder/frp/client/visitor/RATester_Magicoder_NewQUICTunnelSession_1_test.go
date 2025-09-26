package visitor

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewQUICTunnelSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	clientCfg := &v1.ClientCommonConfig{
		ServerAddr: "127.0.0.1",
		ServerPort: 7000,
	}

	session := NewQUICTunnelSession(clientCfg)

	if session == nil {
		t.Error("NewQUICTunnelSession returned nil")
	}

	_, ok := session.(TunnelSession)
	if !ok {
		t.Error("NewQUICTunnelSession did not return a TunnelSession")
	}
}
