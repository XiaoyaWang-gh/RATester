package logs

import (
	"fmt"
	"testing"
)

func TestWriteMsg_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connWriter{}
	lm := &LogMsg{}
	err := c.WriteMsg(lm)
	if err != nil {
		t.Errorf("WriteMsg() error = %v, want nil", err)
	}
}
