package logs

import (
	"fmt"
	"testing"
)

func TestAsync_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	bl.Async(100)
	if !bl.asynchronous {
		t.Error("Expected Async to set asynchronous to true")
	}
	if bl.msgChanLen != 100 {
		t.Error("Expected Async to set msgChanLen to 100")
	}
	if bl.msgChan == nil {
		t.Error("Expected Async to create a msgChan")
	}
	if logMsgPool == nil {
		t.Error("Expected Async to create a logMsgPool")
	}
}
