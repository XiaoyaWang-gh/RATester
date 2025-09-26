package transport

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestRegisterMsgChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	impl := &transporterImpl{
		registry: make(map[string]map[string]chan msg.Message),
	}

	recvCh := make(chan msg.Message)
	laneKey := "lane1"
	msgType := "type1"

	unregister := impl.registerMsgChan(recvCh, laneKey, msgType)

	if _, ok := impl.registry[msgType][laneKey]; !ok {
		t.Errorf("Expected channel to be registered in the registry")
	}

	unregister()

	if _, ok := impl.registry[msgType][laneKey]; ok {
		t.Errorf("Expected channel to be unregistered from the registry")
	}
}
