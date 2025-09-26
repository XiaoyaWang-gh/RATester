package file

import (
	"fmt"
	"testing"
)

func TestAdd_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	flow := &Flow{
		ExportFlow: 100,
		InletFlow:  100,
	}
	flow.Add(10, 10)
	if flow.ExportFlow != 110 {
		t.Errorf("flow.ExportFlow = %d, want 110", flow.ExportFlow)
	}
	if flow.InletFlow != 110 {
		t.Errorf("flow.InletFlow = %d, want 110", flow.InletFlow)
	}
}
