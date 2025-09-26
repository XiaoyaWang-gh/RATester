package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestDeleteTask_1(t *testing.T) {
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

	manager.AddTask("task1", &Task{Taskname: "task1"})
	manager.AddTask("task2", &Task{Taskname: "task2"})

	if len(manager.adminTaskList) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(manager.adminTaskList))
	}

	manager.DeleteTask("task1")

	if len(manager.adminTaskList) != 1 {
		t.Errorf("Expected 1 task after deletion, got %d", len(manager.adminTaskList))
	}

	if _, ok := manager.adminTaskList["task1"]; ok {
		t.Errorf("Task 'task1' should have been deleted")
	}
}
