package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetNext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	task := &Task{
		Next: time.Now().Add(time.Hour),
	}
	next := task.GetNext(ctx)
	if next.IsZero() {
		t.Error("Expected non-zero time, got zero")
	}
}
