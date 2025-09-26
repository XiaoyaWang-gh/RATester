package msg

import (
	"fmt"
	"testing"
)

func TestRun_10(t *testing.T) {
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
	d.Run()
}
