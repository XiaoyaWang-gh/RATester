package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formatter := &PatternLogFormatter{Pattern: "%F:%n|%w %t>> %m", WhenFormat: "2006-01-02"}
	logMsg := &LogMsg{
		Level:      LevelInfo,
		Msg:        "This is a test message",
		When:       time.Now(),
		FilePath:   "test.go",
		LineNumber: 10,
	}
	expected := formatter.ToString(logMsg)
	actual := formatter.Format(logMsg)
	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
