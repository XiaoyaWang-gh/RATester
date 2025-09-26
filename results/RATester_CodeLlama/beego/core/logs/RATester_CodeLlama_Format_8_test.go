package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &consoleWriter{}
	lm := &LogMsg{}
	lm.Level = 1
	lm.Msg = "test"
	lm.When = time.Now()
	lm.FilePath = "test.go"
	lm.LineNumber = 1
	lm.Args = []interface{}{"test"}
	lm.Prefix = "test"
	lm.enableFullFilePath = true
	lm.enableFuncCallDepth = true
	if c.Format(lm) != "test" {
		t.Errorf("consoleWriter.Format() = %v, want %v", c.Format(lm), "test")
	}
}
