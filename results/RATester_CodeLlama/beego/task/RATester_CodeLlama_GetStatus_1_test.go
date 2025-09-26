package task

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var t1 *Task
	if t1.GetStatus(context.Background()) != "" {
		t.Errorf("TestGetStatus failed")
	}
	t1 = &Task{
		Taskname: "test",
		Spec:     &Schedule{},
		SpecStr:  "0 0 0 0 0",
		DoFunc:   nil,
		Prev:     time.Now(),
		Next:     time.Now(),
		Timeout:  0,
		Errlist:  nil,
		ErrLimit: 0,
		errCnt:   0,
	}
	if t1.GetStatus(context.Background()) != "" {
		t.Errorf("TestGetStatus failed")
	}
	t1.Errlist = append(t1.Errlist, &taskerr{
		t:       time.Now(),
		errinfo: "test",
	})
	if t1.GetStatus(context.Background()) != "0001-01-01 00:00:00 +0000 UTC:test<br>" {
		t.Errorf("TestGetStatus failed")
	}
}
