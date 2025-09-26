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

	id := 1
	db.JsonDb.Tasks.Store(id, "task")

	err := db.DelTask(id)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := db.JsonDb.Tasks.Load(id)
	if ok {
		t.Errorf("Expected task to be deleted, but it still exists")
	}
}
