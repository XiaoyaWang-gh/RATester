package task

import (
	"fmt"
	"testing"
)

func TestLen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &MapSorter{
		Keys: []string{"a", "b", "c"},
		Vals: []Tasker{
			&Task{Taskname: "a"},
			&Task{Taskname: "b"},
			&Task{Taskname: "c"},
		},
	}
	if ms.Len() != 3 {
		t.Errorf("ms.Len() = %d, want 3", ms.Len())
	}
}
