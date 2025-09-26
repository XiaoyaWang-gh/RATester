package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewXTCPProxy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{}
	cfg := &v1.XTCPProxyConfig{}
	proxy := NewXTCPProxy(baseProxy, cfg)
	if proxy == nil {
		t.Fatal("NewXTCPProxy failed")
	}
}
