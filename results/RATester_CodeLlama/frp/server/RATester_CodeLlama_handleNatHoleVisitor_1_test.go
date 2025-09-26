package server

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestHandleNatHoleVisitor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	m := &msg.NatHoleVisitor{}
	ctl.handleNatHoleVisitor(m)
}
