package runtime

import (
	"fmt"
	"sync"
	"testing"
)

func TestUpdateServerStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serviceInfo := &ServiceInfo{
		serverStatusMu: sync.RWMutex{},
		serverStatus:   make(map[string]string),
	}

	server := "server1"
	status := "status1"

	serviceInfo.UpdateServerStatus(server, status)

	if serviceInfo.serverStatus[server] != status {
		t.Errorf("Expected status %s for server %s, but got %s", status, server, serviceInfo.serverStatus[server])
	}
}
