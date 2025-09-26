package server

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

	cm := &ControlManager{
		ctlsByRunID: make(map[string]*Control),
	}
	ctl := &Control{}
	old := cm.Add("runID", ctl)
	if old != nil {
		t.Errorf("old is not nil")
	}
	if cm.ctlsByRunID["runID"] != ctl {
		t.Errorf("ctl is not added")
	}
}
