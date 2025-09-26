package msg

import (
	"fmt"
	"testing"
)

func TestDone_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatcher{
		rw:             nil,
		sendCh:         nil,
		doneCh:         nil,
		msgHandlers:    nil,
		defaultHandler: nil,
	}
	if d.Done() != nil {
		t.Error("TestDone failed")
	}
}
