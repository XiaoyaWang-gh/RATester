package task

import (
	"fmt"
	"testing"
)

func TestClearTask_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &taskManager{}
	m.ClearTask()
}
