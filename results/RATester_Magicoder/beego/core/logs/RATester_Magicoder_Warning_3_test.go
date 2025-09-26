package logs

import (
	"fmt"
	"testing"
)

func TestWarning_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: LevelWarn,
	}

	bl.Warning("test warning")

	if bl.level != LevelWarn {
		t.Errorf("Expected level to be %d, but got %d", LevelWarn, bl.level)
	}
}
