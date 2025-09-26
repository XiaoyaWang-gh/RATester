package logs

import (
	"fmt"
	"testing"
)

func Testflush_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		asynchronous: true,
		msgChan:      make(chan *LogMsg, 10),
		outputs:      []*nameLogger{},
	}

	bl.msgChan <- &LogMsg{
		Level: 1,
		Msg:   "test",
	}

	bl.flush()

	if len(bl.msgChan) != 0 {
		t.Errorf("Expected msgChan to be empty after flush, but it's not.")
	}
}
