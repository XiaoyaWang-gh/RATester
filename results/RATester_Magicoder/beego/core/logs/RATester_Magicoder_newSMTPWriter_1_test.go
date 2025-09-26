package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestnewSMTPWriter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	writer := newSMTPWriter()
	if writer == nil {
		t.Error("Expected a non-nil writer, but got nil")
	}

	config := "smtp://user:pass@localhost:25"
	err := writer.Init(config)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	lm := &LogMsg{
		Level:               LevelTrace,
		Msg:                 "Test message",
		When:                time.Now(),
		FilePath:            "test.go",
		LineNumber:          100,
		Args:                []interface{}{"arg1", "arg2"},
		Prefix:              "PREFIX",
		enableFullFilePath:  true,
		enableFuncCallDepth: true,
	}

	err = writer.WriteMsg(lm)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	writer.Destroy()
	writer.Flush()
}
