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
		t.Errorf("Expected no error, got %v", err)
	}
	if result != task {
		t.Errorf("Expected %v, got %v", task, result)
	}

	// Test case 2: Task does not exist
	_, err = db.GetTask(2)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "not found" {
		t.Errorf("Expected 'not found', got %v", err.Error())
	}
}
