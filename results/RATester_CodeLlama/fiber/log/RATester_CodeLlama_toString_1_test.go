package log

import (
	"fmt"
	"testing"
)

func TestToString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lv := Level(1)
	if lv.toString() != "[?1] " {
		t.Errorf("lv.toString() = %v, want [?1] ", lv.toString())
	}
}
