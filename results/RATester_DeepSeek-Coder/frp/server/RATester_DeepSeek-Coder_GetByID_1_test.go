package server

import (
	"fmt"
	"testing"
)

func TestGetByID_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cm := &ControlManager{
		ctlsByRunID: make(map[string]*Control),
	}

	ctl := &Control{
		runID: "test_run_id",
	}

	cm.ctlsByRunID["test_run_id"] = ctl

	gotCtl, gotOk := cm.GetByID("test_run_id")
	if !gotOk {
		t.Errorf("Expected ok to be true, got false")
	}
	if gotCtl != ctl {
		t.Errorf("Expected ctl to be %v, got %v", ctl, gotCtl)
	}

	gotCtl, gotOk = cm.GetByID("non_existent_run_id")
	if gotOk {
		t.Errorf("Expected ok to be false, got true")
	}
	if gotCtl != nil {
		t.Errorf("Expected ctl to be nil, got %v", gotCtl)
	}
}
