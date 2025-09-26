package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestdailyRotate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		dailyOpenTime: time.Now(),
	}

	w.dailyRotate(w.dailyOpenTime)

	if w.needRotateDaily(time.Now().Day()) {
		t.Errorf("Expected needRotateDaily to return false, but got true")
	}
}
