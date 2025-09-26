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

	port := 8080
	ctx := &PortCtx{
		ProxyName:  "test",
		Port:       port,
		Closed:     false,
		UpdateTime: time.Now(),
	}
	pm.usedPorts[port] = ctx

	pm.Release(port)

	if _, ok := pm.usedPorts[port]; ok {
		t.Errorf("Port %d should have been released", port)
	}
	if _, ok := pm.freePorts[port]; !ok {
		t.Errorf("Port %d should have been added to freePorts", port)
	}
	if ctx.Closed != true {
		t.Errorf("PortCtx.Closed should have been set to true")
	}
	if ctx.UpdateTime.Before(time.Now()) {
		t.Errorf("PortCtx.UpdateTime should have been updated")
	}
}
