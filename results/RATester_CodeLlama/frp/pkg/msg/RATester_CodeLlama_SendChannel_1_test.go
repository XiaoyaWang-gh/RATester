package msg

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSendChannel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatcher{
		rw:             nil,
		sendCh:         make(chan Message, 10),
		doneCh:         make(chan struct{}),
		msgHandlers:    make(map[reflect.Type]func(Message)),
		defaultHandler: nil,
	}
	ch := d.SendChannel()
	if ch == nil {
		t.Error("SendChannel() should not return nil")
	}
}
