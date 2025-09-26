package task

import (
	"fmt"
	"testing"
)

func Testinit_19(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	globalTaskManager = newTaskManager()
	if globalTaskManager == nil {
		t.Error("Expected a non-nil task manager, but got nil")
	}
}
