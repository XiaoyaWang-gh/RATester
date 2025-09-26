package logs

import (
	"fmt"
	"testing"
)

func TestAlert_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &BeeLogger{
		level: LevelAlert,
	}

	msg := "Test alert message"
	v := []interface{}{}

	logger.Alert(msg, v...)

	// Add assertions here to verify the behavior of the Alert method
}
