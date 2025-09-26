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

	ctx := context.Background()
	task := &Task{
		Errlist: []*taskerr{
			{
				t:       time.Now(),
				errinfo: "error 1",
			},
			{
				t:       time.Now(),
				errinfo: "error 2",
			},
		},
	}
	expected := task.Errlist[0].t.String() + ":" + task.Errlist[0].errinfo + "<br>" +
		task.Errlist[1].t.String() + ":" + task.Errlist[1].errinfo + "<br>"
	result := task.GetStatus(ctx)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
