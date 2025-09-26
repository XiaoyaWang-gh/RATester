package client

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestHandleNewProxyResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	m := &msg.NewProxyResp{
		ProxyName:  "test",
		RemoteAddr: "127.0.0.1:8080",
		Error:      "test error",
	}
	ctl.handleNewProxyResp(m)
}
