package logs

import (
	"fmt"
	"testing"
)

func TestSetLevel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bl := &BeeLogger{}
	level := 5
	bl.SetLevel(level)
	if bl.level != level {
		t.Errorf("Expected level to be %d, but got %d", level, bl.level)
	}
}
