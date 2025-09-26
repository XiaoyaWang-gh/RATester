package server

import (
	"fmt"
	"testing"
)

func TestDel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cm := &ControlManager{
		ctlsByRunID: make(map[string]*Control),
	}

	ctl := &Control{}
	runID := "test_run_id"
	cm.ctlsByRunID[runID] = ctl

	cm.Del(runID, ctl)

	if _, ok := cm.ctlsByRunID[runID]; ok {
		t.Errorf("Expected control to be deleted, but it still exists")
	}
}
