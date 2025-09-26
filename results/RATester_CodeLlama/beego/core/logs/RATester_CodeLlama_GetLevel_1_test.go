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

	bl := &BeeLogger{}
	bl.level = 1
	if bl.GetLevel() != 1 {
		t.Errorf("GetLevel() = %v, want %v", bl.GetLevel(), 1)
	}
}
