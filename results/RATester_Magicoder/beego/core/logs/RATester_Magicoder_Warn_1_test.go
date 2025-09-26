package logs

import (
	"fmt"
	"testing"
)

func TestWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelWarn,
	}

	bl.Warn("test warn")

	if bl.level != LevelWarn {
		t.Errorf("Expected level to be %d, but got %d", LevelWarn, bl.level)
	}
}
