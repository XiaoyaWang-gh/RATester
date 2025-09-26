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

	if end.Sub(start) < 1*time.Second {
		t.Errorf("Sleep function did not sleep for the expected duration")
	}
}
