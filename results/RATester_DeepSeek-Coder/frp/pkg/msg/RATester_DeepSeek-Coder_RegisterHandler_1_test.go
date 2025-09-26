package msg

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRegisterHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatcher{
		msgHandlers: make(map[reflect.Type]func(Message)),
	}

	msg := &struct{}{}
	handler := func(m Message) {}

	d.RegisterHandler(msg, handler)

	if _, ok := d.msgHandlers[reflect.TypeOf(msg)]; !ok {
		t.Errorf("Expected handler for message type %T not found", msg)
	}
}
