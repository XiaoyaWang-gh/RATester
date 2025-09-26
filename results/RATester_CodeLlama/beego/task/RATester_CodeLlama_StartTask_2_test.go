package task

import (
	"fmt"
	"testing"
)

func TestStartTask_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &taskManager{}
	m.StartTask()
	if m.started != true {
		t.Errorf("StartTask() failed")
	}
}
