package logs

import (
	"fmt"
	"testing"
	"time"
)

func TesthourlyRotate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		hourlyOpenTime: time.Now(),
	}

	w.hourlyRotate(w.hourlyOpenTime)

	// Add assertions here
}
