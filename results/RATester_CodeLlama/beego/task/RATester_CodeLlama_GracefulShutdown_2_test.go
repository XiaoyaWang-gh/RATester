package task

import (
	"fmt"
	"testing"
	"time"
)

func TestGracefulShutdown_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &taskManager{
		adminTaskList: make(map[string]Tasker),
		stop:          make(chan bool),
		changed:       make(chan bool),
	}
	done := m.GracefulShutdown()
	select {
	case <-done:
		t.Log("GracefulShutdown done")
	case <-time.After(time.Second):
		t.Error("GracefulShutdown timeout")
	}
}
