package task

import (
	"fmt"
	"testing"
)

func TestClearTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &taskManager{
		adminTaskList: map[string]Tasker{
			"task1": &Task{},
			"task2": &Task{},
		},
		started: true,
	}

	manager.ClearTask()

	if len(manager.adminTaskList) != 0 {
		t.Errorf("Expected adminTaskList to be empty after ClearTask, but got %d tasks", len(manager.adminTaskList))
	}

	if manager.started {
		t.Error("Expected started to be false after ClearTask, but it's still true")
	}
}
