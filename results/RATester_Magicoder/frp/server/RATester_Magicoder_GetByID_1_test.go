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

	cm.Add("test_run_id", ctl)

	got, ok := cm.GetByID("test_run_id")
	if !ok {
		t.Errorf("Expected to find control with runID 'test_run_id'")
	}

	if got.runID != "test_run_id" {
		t.Errorf("Expected control with runID 'test_run_id', got %v", got.runID)
	}
}
