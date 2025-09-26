package file

import (
	"fmt"
	"testing"
)

func TestUpdateTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	task := &Tunnel{
		Id:       1,
		Port:     8080,
		ServerIp: "127.0.0.1",
	}

	err := db.UpdateTask(task)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	updatedTask, err := db.GetTask(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if updatedTask.Id != 1 {
		t.Errorf("Expected task id to be 1, got %v", updatedTask.Id)
	}

	if updatedTask.Port != 8080 {
		t.Errorf("Expected task port to be 8080, got %v", updatedTask.Port)
	}

	if updatedTask.ServerIp != "127.0.0.1" {
		t.Errorf("Expected task server ip to be 127.0.0.1, got %v", updatedTask.ServerIp)
	}
}
