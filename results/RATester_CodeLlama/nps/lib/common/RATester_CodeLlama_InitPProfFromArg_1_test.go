package common

import (
	"fmt"
	"testing"
)

func TestInitPProfFromArg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitPProfFromArg("")
	InitPProfFromArg(":1234")
	InitPProfFromArg("127.0.0.1:1234")
	InitPProfFromArg("127.0.0.1:1234,127.0.0.1:1235")
	InitPProfFromArg("127.0.0.1:1234,127.0.0.1:1235,127.0.0.1:1236")
}
