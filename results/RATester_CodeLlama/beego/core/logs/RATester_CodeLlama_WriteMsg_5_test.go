package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &consoleWriter{}
	lm := &LogMsg{}
	err := c.WriteMsg(lm)
	if err != nil {
		t.Errorf("WriteMsg() error = %v", err)
		return
	}
	// TODO: Add test cases.
}
