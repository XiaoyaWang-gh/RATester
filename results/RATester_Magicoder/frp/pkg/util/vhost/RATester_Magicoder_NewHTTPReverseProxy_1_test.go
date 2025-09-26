package vhost

import (
	"fmt"
	"testing"
	"time"
)

func TestNewHTTPReverseProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	option := HTTPReverseProxyOptions{
		ResponseHeaderTimeoutS: 60,
	}
	vhostRouter := &Routers{
		indexByDomain: make(map[string]routerByHTTPUser),
	}
	proxy := NewHTTPReverseProxy(option, vhostRouter)

	if proxy == nil {
		t.Error("NewHTTPReverseProxy failed")
	}

	if proxy.responseHeaderTimeout != 60*time.Second {
		t.Error("responseHeaderTimeout is not set correctly")
	}

	if proxy.vhostRouter != vhostRouter {
		t.Error("vhostRouter is not set correctly")
	}

	if proxy.proxy == nil {
		t.Error("proxy is not set correctly")
	}
}
