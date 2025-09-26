package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	task := &Task{
		Timeout: 10 * time.Second,
	}
	timeout := task.GetTimeout(ctx)
	if timeout != 10*time.Second {
		t.Errorf("Expected timeout to be 10 seconds, but got %v", timeout)
	}
}
