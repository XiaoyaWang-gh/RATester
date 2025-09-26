package logs

import (
	"fmt"
	"testing"
)

func TestEmergency_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelEmergency,
	}

	msg := "test emergency"
	bl.Emergency(msg)

	// Check if the message was logged
	select {
	case lm := <-bl.msgChan:
		if lm.Msg != msg {
			t.Errorf("Expected message '%s', got '%s'", msg, lm.Msg)
		}
		if lm.Level != LevelEmergency {
			t.Errorf("Expected level %d, got %d", LevelEmergency, lm.Level)
		}
	default:
		t.Error("No message was logged")
	}
}
