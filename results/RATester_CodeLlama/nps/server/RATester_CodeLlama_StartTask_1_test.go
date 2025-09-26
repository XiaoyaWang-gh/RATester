package server

import (
	"fmt"
	"testing"
)

func TestStartTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	id := 1
	// when
	err := StartTask(id)
	// then
	if err != nil {
		t.Errorf("StartTask() error = %v", err)
		return
	}
	// TODO: check if the task is added to the task list
	// TODO: check if the task status is updated
}
