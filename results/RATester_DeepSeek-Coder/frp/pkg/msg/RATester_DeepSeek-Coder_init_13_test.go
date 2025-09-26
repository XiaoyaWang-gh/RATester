package msg

import (
	"fmt"
	"testing"

	jsonMsg "github.com/fatedier/golib/msg/json"
)

func TestInit_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msgCtl = jsonMsg.NewMsgCtl()
	for typeByte, msg := range msgTypeMap {
		msgCtl.RegisterMsg(typeByte, msg)
	}

	// Add more test cases here
}
