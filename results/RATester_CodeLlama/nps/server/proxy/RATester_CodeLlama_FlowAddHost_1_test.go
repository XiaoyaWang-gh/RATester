package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestFlowAddHost_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	host := &file.Host{
		Flow: &file.Flow{
			ExportFlow: 0,
			InletFlow:  0,
		},
	}
	s := &BaseServer{}
	s.FlowAddHost(host, 100, 200)
	if host.Flow.ExportFlow != 200 {
		t.Error("FlowAddHost error")
	}
	if host.Flow.InletFlow != 100 {
		t.Error("FlowAddHost error")
	}
}
