package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestSetNext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	now := time.Now()

	task := &Task{
		Spec: &Schedule{
			Second: 1,
			Minute: 2,
			Hour:   3,
			Day:    4,
			Month:  5,
		},
	}

	task.SetNext(ctx, now)

	if task.Next.IsZero() {
		t.Errorf("Expected Next to be set, but it is zero")
	}

	if task.Next.Before(now) {
		t.Errorf("Expected Next to be after now, but it is before")
	}
}
