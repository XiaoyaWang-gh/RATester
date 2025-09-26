package logs

import (
	"fmt"
	"testing"
)

func TestTrace_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelDebug,
	}

	msg := "test message"
	v := []interface{}{}

	bl.Trace(msg, v...)

	// assert that the message was written
	// ...
}
