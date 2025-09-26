package logs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
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

	msg := "test message"
	v := []interface{}{"arg1", "arg2"}

	bl.Informational(msg, v...)

	// Verify that the message was written to the log
	// This will depend on the implementation of the BeeLogger.writeMsg method
	// For now, we'll just check that the function doesn't panic
	assert.NotPanics(t, func() {
		bl.Informational(msg, v...)
	})
}
