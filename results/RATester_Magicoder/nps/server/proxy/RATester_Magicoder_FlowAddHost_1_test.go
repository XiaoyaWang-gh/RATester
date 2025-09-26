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

	server := &BaseServer{}
	host := &file.Host{
		Flow: &file.Flow{
			ExportFlow: 0,
			InletFlow:  0,
		},
	}
	in := int64(100)
	out := int64(50)
	server.FlowAddHost(host, in, out)
	if host.Flow.ExportFlow != out || host.Flow.InletFlow != in {
		t.Errorf("Expected export flow to be %d and inlet flow to be %d, but got %d and %d", out, in, host.Flow.ExportFlow, host.Flow.InletFlow)
	}
}
