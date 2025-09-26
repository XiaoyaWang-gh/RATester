package logs

import (
	"fmt"
	"testing"
)

func TestFormat_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}
	lm := &LogMsg{
		Msg: "test message",
	}
	msg := w.Format(lm)
	if msg != "test message" {
		t.Errorf("Expected 'test message', got '%s'", msg)
	}
}
