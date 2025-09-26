package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewTCPMuxProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{
		configurer: &v1.TCPMuxProxyConfig{},
	}

	proxy := NewTCPMuxProxy(baseProxy)

	if proxy == nil {
		t.Error("NewTCPMuxProxy should not return nil")
	}

	_, ok := proxy.(*TCPMuxProxy)
	if !ok {
		t.Error("NewTCPMuxProxy should return a TCPMuxProxy")
	}
}
