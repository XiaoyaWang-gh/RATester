package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestError_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelError,
	}

	msg := "test error message"
	v := []interface{}{"arg1", "arg2"}

	bl.Error(msg, v...)

	// Assuming writeMsg is a mock function that records the message and arguments
	expectedMsg := &LogMsg{
		Level: LevelError,
		Msg:   msg,
		Args:  v,
	}

	if !reflect.DeepEqual(expectedMsg, bl.msgChan) {
		t.Errorf("Expected message to be %v, but got %v", expectedMsg, bl.msgChan)
	}
}
