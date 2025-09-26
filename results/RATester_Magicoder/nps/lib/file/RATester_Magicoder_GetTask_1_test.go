package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{
			Tasks: sync.Map{},
		},
	}

	// Test case 1: Task exists
	task := &Tunnel{Id: 1}
	db.JsonDb.Tasks.Store(1, task)
	result, err := db.GetTask(1)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result.Id != 1 {
		t.Errorf("Expected task id 1, but got %v", result.Id)
	}

	// Test case 2: Task does not exist
	_, err = db.GetTask(2)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
