package alils

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/logs"
)

func TestWriteMsg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &aliLSWriter{}
	lm := &logs.LogMsg{}
	err := c.WriteMsg(lm)
	if err != nil {
		t.Errorf("TestWriteMsg error = %v", err)
	}
}
