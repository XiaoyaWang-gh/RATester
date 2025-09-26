package task

import (
	"fmt"
	"testing"
	"time"
)

func TestStopTask_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	manager := &taskManager{
		adminTaskList: make(map[string]Tasker),
		stop:          make(chan bool),
	}

	manager.StopTask()

	select {
	case <-manager.stop:
		// Test passed
		return
	case <-time.After(time.Second):
		// Test failed
		t.Error("StopTask did not stop the task manager")
	}
}
