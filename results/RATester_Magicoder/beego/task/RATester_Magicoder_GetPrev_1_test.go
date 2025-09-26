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

	task := &Task{
		Prev: time.Now(),
	}
	ctx := context.Background()
	prev := task.GetPrev(ctx)
	if prev.IsZero() {
		t.Error("Expected non-zero time, got zero time")
	}
}
