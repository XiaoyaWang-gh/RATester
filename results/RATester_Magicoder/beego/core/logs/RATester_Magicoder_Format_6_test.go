package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &multiFileLogWriter{}
	lm := &LogMsg{
		Level:      LevelDebug,
		Msg:        "test message",
		When:       time.Now(),
		FilePath:   "test.go",
		LineNumber: 10,
	}
	expected := fmt.Sprintf("%s %s:%d %s", lm.When.Format("2006/01/02 15:04:05"), lm.FilePath, lm.LineNumber, lm.Msg)
	result := w.Format(lm)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
