package logs

import (
	"fmt"
	"testing"
)

func TestGetLevel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{
		level: 5,
	}

	level := bl.GetLevel()

	if level != 5 {
		t.Errorf("Expected level to be 5, but got %d", level)
	}
}
