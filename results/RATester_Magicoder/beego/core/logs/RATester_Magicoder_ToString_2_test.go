package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestToString_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formatter := &PatternLogFormatter{Pattern: "%F:%n|%w %t>> %m", WhenFormat: "2006-01-02"}
	msg := &LogMsg{
		Level:      1,
		Msg:        "test message",
		When:       time.Now(),
		FilePath:   "/path/to/file.go",
		LineNumber: 10,
		Args:       []interface{}{},
	}
	expected := "/path/to/file.go:10|2006-01-02 test message"
	result := formatter.ToString(msg)
	if result != expected {
		t.Errorf("Expected: %s, but got: %s", expected, result)
	}
}
