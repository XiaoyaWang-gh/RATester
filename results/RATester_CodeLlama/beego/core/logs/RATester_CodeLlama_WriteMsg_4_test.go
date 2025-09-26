package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteMsg_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &multiFileLogWriter{}
	lm := &LogMsg{}
	lm.Level = LevelDebug
	lm.Msg = "test"
	lm.When = time.Now()
	lm.FilePath = "test"
	lm.LineNumber = 1
	lm.Args = []interface{}{"test"}
	lm.Prefix = "test"
	lm.enableFullFilePath = true
	lm.enableFuncCallDepth = true
	f.WriteMsg(lm)
}
