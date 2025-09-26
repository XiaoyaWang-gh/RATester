package logs

import (
	"fmt"
	"testing"
)

func TestCritical_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	logger := &BeeLogger{
		level: LevelCritical,
	}

	msg := "test critical message"
	v := []interface{}{}

	logger.Critical(msg, v...)

	// Add your assertions here
}
