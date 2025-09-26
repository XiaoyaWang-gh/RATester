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
			Minute: 1,
			Hour:   1,
			Day:    1,
			Month:  1,
		},
	}
	task.SetNext(ctx, now)

	// Assert that the next time is in the future
	if task.Next.Before(now) {
		t.Errorf("Expected next time to be in the future, but it was not")
	}
}
