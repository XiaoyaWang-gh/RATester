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

	task1 := &Task{
		Taskname: "task1",
		Spec:     &Schedule{Second: 1},
		SpecStr:  "* * * * * *",
	}
	task2 := &Task{
		Taskname: "task2",
		Spec:     &Schedule{Second: 2},
		SpecStr:  "* * * * * *",
	}
	task3 := &Task{
		Taskname: "task3",
		Spec:     &Schedule{Second: 3},
		SpecStr:  "* * * * * *",
	}

	m := map[string]Tasker{
		"task1": task1,
		"task2": task2,
		"task3": task3,
	}

	ms := NewMapSorter(m)

	if len(ms.Keys) != len(m) {
		t.Errorf("Expected %d keys, got %d", len(m), len(ms.Keys))
	}

	if len(ms.Vals) != len(m) {
		t.Errorf("Expected %d values, got %d", len(m), len(ms.Vals))
	}

	for i, key := range ms.Keys {
		if key != ms.Vals[i].(*Task).Taskname {
			t.Errorf("Expected key %s, got %s", key, ms.Vals[i].(*Task).Taskname)
		}
	}
}
