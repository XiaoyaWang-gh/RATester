package ports

import (
	"fmt"
	"testing"
	"time"
)

func TestRelease_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		reservedPorts: make(map[string]*PortCtx),
		usedPorts:     make(map[int]*PortCtx),
		freePorts:     make(map[int]struct{}),
		bindAddr:      "127.0.0.1",
		netType:       "tcp",
	}
	pm.usedPorts[10000] = &PortCtx{
		ProxyName:  "test",
		Port:       10000,
		Closed:     false,
		UpdateTime: time.Now(),
	}
	pm.Release(10000)
	if _, ok := pm.usedPorts[10000]; ok {
		t.Error("release failed")
	}
}
