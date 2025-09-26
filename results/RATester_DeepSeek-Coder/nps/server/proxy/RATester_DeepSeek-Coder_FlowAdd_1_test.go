package proxy

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestFlowAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &BaseServer{
		id: 1,
		task: &file.Tunnel{
			Flow: &file.Flow{
				ExportFlow: 0,
				InletFlow:  0,
			},
		},
	}

	server.FlowAdd(10, 20)

	if server.task.Flow.ExportFlow != 20 {
		t.Errorf("Expected ExportFlow to be 20, got %d", server.task.Flow.ExportFlow)
	}

	if server.task.Flow.InletFlow != 10 {
		t.Errorf("Expected InletFlow to be 10, got %d", server.task.Flow.InletFlow)
	}
}
