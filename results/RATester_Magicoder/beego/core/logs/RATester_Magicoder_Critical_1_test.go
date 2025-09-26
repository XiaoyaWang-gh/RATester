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

	bl := &BeeLogger{
		level: LevelCritical,
	}

	msg := "test critical"
	bl.Critical(msg)

	if bl.level != LevelCritical {
		t.Errorf("Expected level to be %d, but got %d", LevelCritical, bl.level)
	}
}
