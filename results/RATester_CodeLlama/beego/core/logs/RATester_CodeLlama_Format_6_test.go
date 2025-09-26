package logs

import (
	"fmt"
	"testing"
)

func TestFormat_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lm := &LogMsg{
		Level: LevelDebug,
		Msg:   "test",
	}
	if got := lm.OldStyleFormat(); got != "DEBUG test" {
		t.Errorf("LogMsg.OldStyleFormat() = %v, want %v", got, "DEBUG test")
	}
}
