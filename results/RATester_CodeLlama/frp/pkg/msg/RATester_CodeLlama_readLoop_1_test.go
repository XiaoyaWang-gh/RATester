package msg

import (
	"fmt"
	"testing"
)

func TestReadLoop_1(t *testing.T) {
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
	d.readLoop()
}
