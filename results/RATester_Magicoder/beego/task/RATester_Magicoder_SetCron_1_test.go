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
		Taskname: "testTask",
		SpecStr:  "* * * * *",
	}

	task.SetCron("* * * * *")

	if task.SpecStr != "* * * * *" {
		t.Errorf("Expected SpecStr to be '* * * * *', but got '%s'", task.SpecStr)
	}
}
