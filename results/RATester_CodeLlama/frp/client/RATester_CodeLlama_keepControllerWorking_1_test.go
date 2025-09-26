package client

import (
	"fmt"
	"testing"
	"time"
)

func TestKeepControllerWorking_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	svr.ctl = &Control{}
	svr.ctl.Done()
	svr.loopLoginUntilSuccess(20*time.Second, false)
	if svr.ctl != nil {
		<-svr.ctl.Done()
		t.Error("control is closed and try another loop")
	}
}
