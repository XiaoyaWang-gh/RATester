package logs

import (
	"fmt"
	"testing"
)

func TestEmergency_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelEmergency,
	}

	msg := "test emergency"
	bl.Emergency(msg)

	if bl.level != LevelEmergency {
		t.Errorf("Expected level to be %d, but got %d", LevelEmergency, bl.level)
	}
}
