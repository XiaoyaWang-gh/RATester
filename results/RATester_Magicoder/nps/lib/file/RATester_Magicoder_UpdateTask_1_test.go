package file

import (
	"fmt"
	"sync"
	"testing"
)

func TestUpdateTask_1(t *testing.T) {
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

	task := &Tunnel{
		Id: 1,
		// other fields
	}

	err := db.UpdateTask(task)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, ok := db.JsonDb.Tasks.Load(task.Id)
	if !ok {
		t.Errorf("Expected task to be stored in the map, but it was not")
	}
}
