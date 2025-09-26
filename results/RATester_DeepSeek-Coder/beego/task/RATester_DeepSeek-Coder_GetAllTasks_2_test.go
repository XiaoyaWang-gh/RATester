package task

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetAllTasks_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	taskManager := &taskManager{
		adminTaskList: make(map[string]Tasker),
		taskLock:      sync.RWMutex{},
	}

	task1 := &Task{}
	task2 := &Task{}

	taskManager.AddTask("task1", task1)
	taskManager.AddTask("task2", task2)

	tasks := taskManager.GetAllTasks()

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if tasks[0] != task1 || tasks[1] != task2 {
		t.Errorf("Expected tasks in correct order")
	}
}
