package server

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
	"github.com/fatedier/frp/server/controller"
)

func TestHandleNatHoleReport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{
		rc: &controller.ResourceController{},
	}

	testMsg := &msg.NatHoleReport{
		// fill in the fields of the test message
	}

	ctl.handleNatHoleReport(testMsg)

	// assert the expected results
}
