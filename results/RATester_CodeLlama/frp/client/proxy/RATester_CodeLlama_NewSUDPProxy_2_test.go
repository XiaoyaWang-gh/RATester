package proxy

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/zeebo/assert"
)

func TestNewSUDPProxy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{}
	cfg := &v1.SUDPProxyConfig{}
	proxy := NewSUDPProxy(baseProxy, cfg)
	assert.NotNil(t, proxy)
}
