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
		Keys: []string{"key1", "key2", "key3"},
		Vals: []Tasker{&Task{}, &Task{}, &Task{}},
	}

	if ms.Len() != 3 {
		t.Errorf("Expected length of 3, got %d", ms.Len())
	}
}
