package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewHTTPSProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{
		name:       "test",
		configurer: &v1.HTTPSProxyConfig{},
	}

	proxy := NewHTTPSProxy(baseProxy)

	if proxy == nil {
		t.Error("Expected a non-nil proxy, but got nil")
	}

	if proxy.GetName() != "test" {
		t.Errorf("Expected proxy name to be 'test', but got '%s'", proxy.GetName())
	}

	if _, ok := proxy.GetConfigurer().(*v1.HTTPSProxyConfig); !ok {
		t.Error("Expected a HTTPSProxyConfig configurer, but got a different type")
	}
}
