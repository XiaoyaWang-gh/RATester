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

	ctx := context.Background()
	now := time.Now()
	task := &Task{
		Prev: time.Time{},
	}
	task.SetPrev(ctx, now)
	if task.Prev != now {
		t.Errorf("Expected Prev to be %v, but got %v", now, task.Prev)
	}
}
