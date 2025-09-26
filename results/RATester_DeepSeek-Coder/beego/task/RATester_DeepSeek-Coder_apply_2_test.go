package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestApply_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	task := &Task{
		Taskname: "test",
		Spec:     &Schedule{},
		SpecStr:  "* * * * *",
		DoFunc:   func(context.Context) error { return nil },
		Prev:     time.Now(),
		Next:     time.Now(),
		Timeout:  10 * time.Second,
	}

	optionFunc(func(t *Task) {
		t.Taskname = "test2"
	}).apply(task)

	if task.Taskname != "test2" {
		t.Errorf("Expected Taskname to be 'test2', got %s", task.Taskname)
	}
}
