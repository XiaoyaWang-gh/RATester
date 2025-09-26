package task

import (
	"fmt"
	"testing"
)

func TestnewTaskManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := newTaskManager()

	if manager == nil {
		t.Error("Expected a non-nil taskManager, but got nil")
	}

	if manager.adminTaskList == nil {
		t.Error("Expected a non-nil adminTaskList, but got nil")
	}

	if manager.stop == nil {
		t.Error("Expected a non-nil stop channel, but got nil")
	}

	if manager.changed == nil {
		t.Error("Expected a non-nil changed channel, but got nil")
	}

	if manager.started != false {
		t.Error("Expected started to be false, but got true")
	}
}
