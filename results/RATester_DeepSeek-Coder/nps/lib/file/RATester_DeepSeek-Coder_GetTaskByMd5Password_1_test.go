package file

import (
	"fmt"
	"sync"
	"testing"

	"ehang.io/nps/lib/crypt"
)

func TestGetTaskByMd5Password_1(t *testing.T) {
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

	// Add a task
	task := &Tunnel{
		Id:       1,
		Password: "password",
	}
	db.JsonDb.Tasks.Store(task.Id, task)

	// Test GetTaskByMd5Password
	md5Password := crypt.Md5("password")
	result := db.GetTaskByMd5Password(md5Password)
	if result == nil {
		t.Errorf("Expected task, got nil")
	}
	if result.Id != task.Id {
		t.Errorf("Expected task id %d, got %d", task.Id, result.Id)
	}

	// Test GetTaskByMd5Password with non-existing password
	nonExistingPassword := crypt.Md5("non-existing-password")
	result = db.GetTaskByMd5Password(nonExistingPassword)
	if result != nil {
		t.Errorf("Expected nil, got task")
	}
}
