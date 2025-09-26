package task

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeoutOption_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	timeout := time.Duration(10)
	opt := TimeoutOption(timeout)
	task := &Task{}
	opt.apply(task)
	if task.Timeout != timeout {
		t.Errorf("TimeoutOption failed")
	}
}
