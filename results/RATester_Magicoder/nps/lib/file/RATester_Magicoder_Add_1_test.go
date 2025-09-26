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
		ExportFlow: 0,
		InletFlow:  0,
		FlowLimit:  100,
	}

	flow.Add(50, 20)

	if flow.ExportFlow != 20 {
		t.Errorf("Expected ExportFlow to be 20, but got %d", flow.ExportFlow)
	}

	if flow.InletFlow != 50 {
		t.Errorf("Expected InletFlow to be 50, but got %d", flow.InletFlow)
	}
}
