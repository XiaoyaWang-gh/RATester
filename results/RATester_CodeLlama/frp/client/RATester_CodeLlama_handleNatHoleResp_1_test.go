package client

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestHandleNatHoleResp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	m := &msg.NatHoleResp{}
	ctl.handleNatHoleResp(m)
}
