package task

import (
	"context"
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
		Keys: []string{"key1", "key2", "key3"},
		Vals: []Tasker{&Task{}, &Task{}, &Task{}},
	}

	// Test case 1: Task with non-zero next time is less than Task with zero next time
	ms.Vals[0].SetNext(context.Background(), time.Now().Add(1*time.Hour))
	ms.Vals[1].SetNext(context.Background(), time.Time{})
	if !ms.Less(0, 1) {
		t.Errorf("Expected Task with non-zero next time to be less than Task with zero next time")
	}

	// Test case 2: Task with earlier next time is less than Task with later next time
	ms.Vals[0].SetNext(context.Background(), time.Now().Add(1*time.Hour))
	ms.Vals[1].SetNext(context.Background(), time.Now().Add(2*time.Hour))
	if !ms.Less(0, 1) {
		t.Errorf("Expected Task with earlier next time to be less than Task with later next time")
	}

	// Test case 3: Task with earlier next time is not less than Task with earlier next time
	ms.Vals[0].SetNext(context.Background(), time.Now().Add(1*time.Hour))
	ms.Vals[1].SetNext(context.Background(), time.Now().Add(1*time.Hour))
	if ms.Less(0, 1) {
		t.Errorf("Expected Task with same next time not to be less than Task with same next time")
	}
}
