package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connWriter{}
	lm := &LogMsg{}
	lm.Msg = "test"
	lm.Level = 1
	lm.When = time.Now()
	lm.FilePath = "test.go"
	lm.LineNumber = 1
	lm.Args = []interface{}{"test"}
	lm.Prefix = "test"
	lm.enableFullFilePath = true
	lm.enableFuncCallDepth = true
	if c.Format(lm) != lm.OldStyleFormat() {
		t.Errorf("connWriter.Format() = %v, want %v", c.Format(lm), lm.OldStyleFormat())
	}
}
