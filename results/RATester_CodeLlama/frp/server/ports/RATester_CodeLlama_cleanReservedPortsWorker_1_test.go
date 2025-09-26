package ports

import (
	"fmt"
	"testing"
	"time"
)

func TestCleanReservedPortsWorker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		reservedPorts: map[string]*PortCtx{
			"test": {
				ProxyName:  "test",
				Port:       10000,
				Closed:     false,
				UpdateTime: time.Now(),
			},
		},
	}
	go pm.cleanReservedPortsWorker()
	time.Sleep(time.Second)
	pm.mu.Lock()
	if _, ok := pm.reservedPorts["test"]; ok {
		t.Error("cleanReservedPortsWorker failed")
	}
	pm.mu.Unlock()
}
