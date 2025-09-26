package task

import (
	"fmt"
	"testing"
)

func TestNewTaskManager_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := newTaskManager()

	if manager == nil {
		t.Errorf("Expected a non-nil taskManager, got nil")
	}

	if manager.adminTaskList == nil {
		t.Errorf("Expected a non-nil adminTaskList, got nil")
	}

	if manager.stop == nil {
		t.Errorf("Expected a non-nil stop channel, got nil")
	}

	if manager.changed == nil {
		t.Errorf("Expected a non-nil changed channel, got nil")
	}

	if manager.started != false {
		t.Errorf("Expected started to be false, got %v", manager.started)
	}
}
