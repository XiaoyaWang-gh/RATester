package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSetPrev_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	task := &Task{
		Prev: time.Time{},
	}

	now := time.Now()
	task.SetPrev(ctx, now)

	if task.Prev != now {
		t.Errorf("Expected Prev to be %v, but got %v", now, task.Prev)
	}
}
