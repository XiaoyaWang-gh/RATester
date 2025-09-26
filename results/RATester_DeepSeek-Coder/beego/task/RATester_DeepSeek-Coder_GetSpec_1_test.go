package task

import (
	"context"
	"fmt"
	"testing"
)

func TestGetSpec_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	task := &Task{
		SpecStr: "* * * * *",
	}
	spec := task.GetSpec(ctx)
	if spec != "* * * * *" {
		t.Errorf("Expected '* * * * *', got '%s'", spec)
	}
}
