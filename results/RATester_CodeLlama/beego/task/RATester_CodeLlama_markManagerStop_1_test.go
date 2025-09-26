package task

import (
	"fmt"
	"testing"
)

func TestMarkManagerStop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &taskManager{}
	m.markManagerStop()
	if m.started {
		t.Errorf("markManagerStop failed")
	}
}
