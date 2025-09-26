package logs

import (
	"fmt"
	"testing"
)

func TestDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelDebug,
	}

	format := "test message"
	v := []interface{}{"arg1", "arg2"}

	bl.Debug(format, v...)

	// Assert that the log message was written
	// ...
}
