package logs

import (
	"fmt"
	"testing"
)

func TestwriteToLoggers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		outputs: []*nameLogger{
			{
				name: "test",
			},
		},
	}

	lm := &LogMsg{
		Msg: "test message",
	}

	bl.writeToLoggers(lm)

	// Assert that the log message was written to the logger
	// You can use the testing package's assert functions here
}
