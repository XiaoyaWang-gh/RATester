package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestMarkManagerStop_1(t *testing.T) {
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
		started:       true,
		wait:          sync.WaitGroup{},
	}

	manager.markManagerStop()

	if manager.started {
		t.Errorf("Expected started to be false, got true")
	}
}
