package task

import (
	"fmt"
	"testing"
	"time"
)

func TestLess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ms := &MapSorter{
		Keys: []string{"key1", "key2"},
		Vals: []Tasker{
			&Task{
				Next: time.Now().Add(time.Hour),
			},
			&Task{
				Next: time.Now().Add(time.Minute),
			},
		},
	}

	if ms.Less(0, 1) {
		t.Errorf("Expected false, got true")
	}

	ms.Vals[0].(*Task).Next = time.Time{}
	if !ms.Less(0, 1) {
		t.Errorf("Expected true, got false")
	}

	ms.Vals[0].(*Task).Next = time.Now().Add(time.Minute)
	ms.Vals[1].(*Task).Next = time.Now().Add(time.Hour)
	if ms.Less(0, 1) {
		t.Errorf("Expected false, got true")
	}
}
