package try

import (
	"fmt"
	"testing"
	"time"
)

func TestSleep_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	start := time.Now()
	Sleep(1 * time.Second)
	end := time.Now()

	diff := end.Sub(start)
	if diff < 1*time.Second || diff > 1*time.Second+100*time.Millisecond {
		t.Errorf("Expected sleep duration to be around 1 second, got %v", diff)
	}
}
