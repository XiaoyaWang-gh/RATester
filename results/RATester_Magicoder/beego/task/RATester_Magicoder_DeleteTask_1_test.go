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
		adminTaskList: map[string]Tasker{
			"task1": &Task{},
			"task2": &Task{},
		},
		taskLock: sync.RWMutex{},
		stop:     make(chan bool),
		changed:  make(chan bool),
		started:  false,
		wait:     sync.WaitGroup{},
	}

	manager.DeleteTask("task1")

	if _, ok := manager.adminTaskList["task1"]; ok {
		t.Error("Expected task1 to be deleted")
	}

	if _, ok := manager.adminTaskList["task2"]; !ok {
		t.Error("Expected task2 to still exist")
	}
}
