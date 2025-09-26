package logs

import (
	"fmt"
	"testing"
)

func TestInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &BeeLogger{
		level: LevelInfo,
	}

	msg := "test message"
	args := []interface{}{"arg1", "arg2"}

	logger.Info(msg, args...)

	// Verify that the message was written to the logger's output
	// This will depend on the specific implementation of the logger's output
	// For example, if the output is a file, you could read the file and verify its content
}
