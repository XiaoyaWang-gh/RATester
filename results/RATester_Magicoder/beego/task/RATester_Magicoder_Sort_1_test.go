package task

import (
	"fmt"
	"testing"
)

func TestSort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &MapSorter{
		Keys: []string{"b", "a", "c"},
		Vals: []Tasker{
			&Task{Taskname: "b"},
			&Task{Taskname: "a"},
			&Task{Taskname: "c"},
		},
	}
	ms.Sort()

	expected := []string{"a", "b", "c"}
	for i, key := range ms.Keys {
		if key != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], key)
		}
	}
}
