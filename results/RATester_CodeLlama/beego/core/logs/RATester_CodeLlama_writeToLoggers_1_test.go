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

	// TODO: setup your test
	lm := &LogMsg{}
	bl := &BeeLogger{}
	// TODO: perform your test
	bl.writeToLoggers(lm)
	// TODO: verify your results
}
