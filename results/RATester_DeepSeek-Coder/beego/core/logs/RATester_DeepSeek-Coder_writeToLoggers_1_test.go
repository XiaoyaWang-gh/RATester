package logs

import (
	"fmt"
	"testing"
)

func TestWriteToLoggers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testLogger := &BeeLogger{
		outputs: []*nameLogger{
			{
				name: "testLogger1",
			},
			{
				name: "testLogger2",
			},
		},
	}

	testMsg := &LogMsg{
		Msg: "test message",
	}

	testLogger.writeToLoggers(testMsg)

	for _, output := range testLogger.outputs {
		if output.name != "testLogger1" && output.name != "testLogger2" {
			t.Errorf("Unexpected logger name: %s", output.name)
		}
	}
}
