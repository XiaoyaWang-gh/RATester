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
	msg := &LogMsg{
		Level:               1,
		Msg:                 "test message",
		When:                time.Now(),
		FilePath:            "test.go",
		LineNumber:          100,
		enableFullFilePath:  true,
		enableFuncCallDepth: true,
	}
	expected := fmt.Sprintf("%s:%s|%s %s>> %s", msg.When.Format("2006-01-02"), "test.go", "INFO", msg.When.Format("15:04:05"), msg.Msg)
	result := formatter.Format(msg)
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
