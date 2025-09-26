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

	m := &taskManager{
		adminTaskList: make(map[string]Tasker),
		stop:          make(chan bool),
	}
	m.StopTask()
	select {
	case <-m.stop:
		t.Log("stop task success")
	default:
		t.Error("stop task failed")
	}
}
