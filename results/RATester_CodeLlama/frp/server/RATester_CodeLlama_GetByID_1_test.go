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
	ctl := &Control{}
	cm.Add("runID", ctl)
	got, ok := cm.GetByID("runID")
	if !ok {
		t.Errorf("GetByID() got = %v, want %v", got, ctl)
	}
}
