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

	// Add a task with a password
	task := &Tunnel{
		Password: "password",
	}
	db.JsonDb.Tasks.Store("task1", task)

	// Test the function
	result := db.GetTaskByMd5Password(crypt.Md5("password"))

	// Verify the result
	if result != task {
		t.Errorf("Expected %v, got %v", task, result)
	}
}
