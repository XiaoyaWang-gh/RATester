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

	m := &taskManager{
		adminTaskList: make(map[string]Tasker),
		taskLock:      sync.RWMutex{},
		stop:          make(chan bool),
		changed:       make(chan bool),
		started:       false,
		wait:          sync.WaitGroup{},
	}

	taskname := "test"
	m.AddTask(taskname, nil)

	m.DeleteTask(taskname)

	if _, ok := m.adminTaskList[taskname]; ok {
		t.Errorf("DeleteTask failed")
	}
}
