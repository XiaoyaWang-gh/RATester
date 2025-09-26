package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestwriteMsg_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		init:                true,
		enableFuncCallDepth: true,
		enableFullFilePath:  true,
		asynchronous:        true,
		logWithNonBlocking:  true,
		level:               LevelEmergency,
		loggerFuncCallDepth: 1,
		prefix:              "test",
		msgChanLen:          100,
		msgChan:             make(chan *LogMsg, 100),
		closeChan:           make(chan struct{}),
		flushChan:           make(chan struct{}),
		outputs:             []*nameLogger{},
	}

	lm := &LogMsg{
		Level:               LevelEmergency,
		Msg:                 "test message",
		When:                time.Now(),
		FilePath:            "test.go",
		LineNumber:          10,
		Args:                []interface{}{"arg1", "arg2"},
		Prefix:              "test",
		enableFullFilePath:  true,
		enableFuncCallDepth: true,
	}

	err := bl.writeMsg(lm)
	if err != nil {
		t.Errorf("writeMsg failed: %v", err)
	}

	select {
	case msg := <-bl.msgChan:
		if msg.Level != lm.Level || msg.Msg != lm.Msg || msg.When.Unix() != lm.When.Unix() || msg.FilePath != lm.FilePath || msg.LineNumber != lm.LineNumber || msg.Prefix != lm.Prefix || msg.enableFullFilePath != lm.enableFullFilePath || msg.enableFuncCallDepth != lm.enableFuncCallDepth {
			t.Errorf("writeMsg failed: received message does not match sent message")
		}
	case <-time.After(time.Second):
		t.Errorf("writeMsg failed: timed out waiting for message")
	}
}
