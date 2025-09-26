package task

import (
	"fmt"
	"testing"
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
		// The stop channel should have a value
		t.Errorf("Expected stop channel to have a value")
	default:
		// The stop channel should be empty
	}
}
