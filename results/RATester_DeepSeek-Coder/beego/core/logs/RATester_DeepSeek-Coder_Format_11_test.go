package logs

import (
	"fmt"
	"testing"
	"time"
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
		Msg:    "test message",
		When:   time.Now(),
		Level:  1,
		Prefix: "test",
	}
	expected := fmt.Sprintf("%s %s\n", lm.When.Format("2006/01/02 15:04:05"), lm.Msg)
	result := w.Format(lm)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
