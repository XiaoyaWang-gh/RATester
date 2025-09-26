package nathole

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/msg"
)

func TestHandleReport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
		sessions:   make(map[string]*Session),
		analyzer:   &Analyzer{records: make(map[string]*MakeHoleRecords)},
	}
	m := &msg.NatHoleReport{
		Sid:     "sid",
		Success: true,
	}
	c.HandleReport(m)
}
