package logs

import (
	"fmt"
	"testing"
)

func TestInformational_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelInfo,
	}

	format := "test message"
	v := []interface{}{"arg1", "arg2"}

	bl.Informational(format, v...)

	// Assert that the log message was written
	// ...
}
