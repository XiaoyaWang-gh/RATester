package debug

import (
	"fmt"
	"testing"
	"time"
)

func TestStop_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testTimer := &timer{
		start: time.Now(),
	}

	testTimer.Stop()

	if testTimer.elapsed == 0 {
		t.Errorf("Expected elapsed time to be greater than 0, but got 0")
	}
}
