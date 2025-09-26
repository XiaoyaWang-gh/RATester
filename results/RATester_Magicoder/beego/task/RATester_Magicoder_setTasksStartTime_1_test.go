package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestsetTasksStartTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	taskManager := &taskManager{
		adminTaskList: map[string]Tasker{
			"task1": &Task{
				SpecStr: "* * * * * *",
			},
			"task2": &Task{
				SpecStr: "* * * * * *",
			},
		},
	}
	taskManager.setTasksStartTime(now)
	for _, task := range taskManager.adminTaskList {
		if task.GetNext(context.Background()).Before(now) {
			t.Errorf("Task next time should be after now")
		}
	}
}
