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

	ctl := &Control{
		// initialize Control fields
	}

	old := cm.Add("runID", ctl)

	if old != nil {
		t.Errorf("Expected nil, got %v", old)
	}

	old = cm.Add("runID", ctl)

	if old == nil {
		t.Errorf("Expected non-nil, got nil")
	}
}
