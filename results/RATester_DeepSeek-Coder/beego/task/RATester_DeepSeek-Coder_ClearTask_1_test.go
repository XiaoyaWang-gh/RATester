package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestClearTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &taskManager{
		adminTaskList: make(map[string]Tasker),
		taskLock:      sync.RWMutex{},
		stop:          make(chan bool),
		changed:       make(chan bool),
		started:       false,
	}

	manager.AddTask("task1", &Task{})
	manager.AddTask("task2", &Task{})

	if len(manager.adminTaskList) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(manager.adminTaskList))
	}

	manager.ClearTask()

	if len(manager.adminTaskList) != 0 {
		t.Errorf("Expected 0 tasks after clear, got %d", len(manager.adminTaskList))
	}

	if manager.started {
		t.Errorf("Expected started to be false after clear")
	}
}
