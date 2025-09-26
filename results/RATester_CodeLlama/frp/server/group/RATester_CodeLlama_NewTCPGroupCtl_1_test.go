package group

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/server/ports"
)

func TestNewTCPGroupCtl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &ports.Manager{}
	tgc := NewTCPGroupCtl(pm)
	if tgc == nil {
		t.Error("NewTCPGroupCtl return nil")
	}
}
