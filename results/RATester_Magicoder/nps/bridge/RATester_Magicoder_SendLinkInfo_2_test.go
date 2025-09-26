package bridge

import (
	"fmt"
	"sync"
	"testing"

	"ehang.io/nps/lib/conn"
	"ehang.io/nps/lib/file"
)

func TestSendLinkInfo_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &Bridge{
		Client:   sync.Map{},
		Register: sync.Map{},
	}

	clientId := 1
	link := &conn.Link{
		Host: "localhost:8080",
	}
	tunnel := &file.Tunnel{}

	target, err := bridge.SendLinkInfo(clientId, link, tunnel)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if target == nil {
		t.Error("Expected target to be not nil")
	}
}
