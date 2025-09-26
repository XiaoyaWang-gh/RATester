package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Rotate:   true,
		Daily:    true,
		Hourly:   true,
		Filename: "test.log",
		MaxLines: 10000,
		MaxFiles: 10,
		MaxSize:  1024,
		MaxDays:  7,
		MaxHours: 24,
		Level:    1,
	}

	lm := &LogMsg{
		Level:      1,
		Msg:        "test message",
		When:       time.Now(),
		FilePath:   "test.go",
		LineNumber: 1,
	}

	err := w.WriteMsg(lm)
	if err != nil {
		t.Errorf("WriteMsg failed: %v", err)
	}
}
