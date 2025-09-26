package task

import (
	"errors"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/admin"
)

func TestExecute_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &runTaskCommand{}

	// Test case 1: task name not passed
	result := r.Execute()
	if result.Status != 400 || result.Error == nil || result.Content != nil {
		t.Errorf("Test case 1 failed. Expected: %v, got: %v", &admin.Result{Status: 400, Error: errors.New("task name not passed")}, result)
	}

	// Test case 2: parameter is invalid
	result = r.Execute(123)
	if result.Status != 400 || result.Error == nil || result.Content != nil {
		t.Errorf("Test case 2 failed. Expected: %v, got: %v", &admin.Result{Status: 400, Error: errors.New("parameter is invalid")}, result)
	}

	// Test case 3: task not found
	result = r.Execute("non-existent-task")
	if result.Status != 400 || result.Error == nil || result.Content != nil {
		t.Errorf("Test case 3 failed. Expected: %v, got: %v", &admin.Result{Status: 400, Error: fmt.Errorf("task with name non-existent-task not found")}, result)
	}

	// Test case 4: task run successfully
	result = r.Execute("existing-task")
	if result.Status != 200 || result.Error != nil || result.Content == nil {
		t.Errorf("Test case 4 failed. Expected: %v, got: %v", &admin.Result{Status: 200, Content: "task-status"}, result)
	}
}
