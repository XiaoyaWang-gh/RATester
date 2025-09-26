package task

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestRun_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	task := &Task{
		DoFunc: func(ctx context.Context) error {
			return errors.New("test error")
		},
		ErrLimit: 1,
		Errlist:  make([]*taskerr, 1),
	}

	err := task.Run(ctx)
	if err == nil {
		t.Error("expected error, got nil")
	}

	if len(task.Errlist) != 1 {
		t.Errorf("expected Errlist length 1, got %d", len(task.Errlist))
	}

	if task.Errlist[0] == nil {
		t.Error("expected non-nil taskerr, got nil")
	}

	if task.Errlist[0].errinfo != "test error" {
		t.Errorf("expected errinfo 'test error', got '%s'", task.Errlist[0].errinfo)
	}
}
