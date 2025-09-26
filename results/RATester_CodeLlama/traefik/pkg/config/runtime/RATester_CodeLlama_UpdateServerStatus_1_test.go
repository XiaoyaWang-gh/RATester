package runtime

import (
	"fmt"
	"testing"
)

func TestUpdateServerStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &ServiceInfo{}
	server := "server"
	status := "status"

	s.UpdateServerStatus(server, status)

	if s.serverStatus == nil {
		t.Errorf("serverStatus should not be nil")
	}

	if s.serverStatus[server] != status {
		t.Errorf("serverStatus[server] should be %s, got %s", status, s.serverStatus[server])
	}
}
