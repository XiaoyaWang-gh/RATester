package proxy

import (
	"fmt"
	"testing"
)

func TestFlowAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &BaseServer{}
	s.FlowAdd(1, 2)
	if s.task.Flow.ExportFlow != 2 {
		t.Error("FlowAdd error")
	}
	if s.task.Flow.InletFlow != 1 {
		t.Error("FlowAdd error")
	}
}
