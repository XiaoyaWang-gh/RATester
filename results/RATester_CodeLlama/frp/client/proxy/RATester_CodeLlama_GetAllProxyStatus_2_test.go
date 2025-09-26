package proxy

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestGetAllProxyStatus_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		proxies: make(map[string]*Wrapper),
	}
	pm.proxies["proxy1"] = &Wrapper{
		WorkingStatus: WorkingStatus{
			Name: "proxy1",
			Type: "http",
		},
	}
	pm.proxies["proxy2"] = &Wrapper{
		WorkingStatus: WorkingStatus{
			Name: "proxy2",
			Type: "tcp",
		},
	}
	ps := pm.GetAllProxyStatus()
	assert.Equal(t, 2, len(ps))
	assert.Equal(t, "proxy1", ps[0].Name)
	assert.Equal(t, "proxy2", ps[1].Name)
}
