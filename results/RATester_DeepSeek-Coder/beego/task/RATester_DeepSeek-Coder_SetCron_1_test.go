package task

import (
	"fmt"
	"testing"
)

func TestSetCron_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	task := &Task{
		Taskname: "TestTask",
		SpecStr:  "* * * * *",
	}

	task.SetCron("0 0 12 * * ?")

	if task.SpecStr != "0 0 12 * * ?" {
		t.Errorf("Expected SpecStr to be '0 0 12 * * ?', but got %s", task.SpecStr)
	}
}
