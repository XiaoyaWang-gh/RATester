package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestStartTask_2(t *testing.T) {
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

	manager.StartTask()

	if !manager.started {
		t.Error("Expected task to be started")
	}
}
