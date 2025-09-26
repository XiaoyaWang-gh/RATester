package transport

import (
	"fmt"
	"testing"

	"github.com/fatedier/golib/msg/json"
)

func TestRegisterMsgChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	impl := &transporterImpl{
		sendCh:   make(chan json.Message),
		registry: make(map[string]map[string]chan json.Message),
	}

	recvCh := make(chan json.Message)
	laneKey := "testLaneKey"
	msgType := "testMsgType"

	unregister := impl.registerMsgChan(recvCh, laneKey, msgType)

	if _, ok := impl.registry[msgType][laneKey]; !ok {
		t.Errorf("Expected to find laneKey in registry")
	}

	unregister()

	if _, ok := impl.registry[msgType][laneKey]; ok {
		t.Errorf("Expected to not find laneKey in registry after unregister")
	}
}
