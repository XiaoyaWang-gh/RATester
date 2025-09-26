package task

import (
	"fmt"
	"testing"
)

func TestNewMapSorter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := make(map[string]Tasker)
	m["task1"] = &Task{Taskname: "task1"}
	m["task2"] = &Task{Taskname: "task2"}
	m["task3"] = &Task{Taskname: "task3"}
	ms := NewMapSorter(m)
	if len(ms.Keys) != len(m) {
		t.Errorf("NewMapSorter failed")
	}
}
