package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetPrev_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	task := &Task{
		Prev: time.Now(),
	}

	prev := task.GetPrev(ctx)

	if prev.IsZero() {
		t.Errorf("Expected non-zero time, got zero time")
	}

	if !prev.Equal(task.Prev) {
		t.Errorf("Expected %v, got %v", task.Prev, prev)
	}
}
