package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestDelTask_1(t *testing.T) {
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

	// Add a task to the database
	db.JsonDb.Tasks.Store(1, "task1")

	// Delete the task
	err := db.DelTask(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the task is deleted
	_, ok := db.JsonDb.Tasks.Load(1)
	if ok {
		t.Errorf("Expected task to be deleted")
	}
}
