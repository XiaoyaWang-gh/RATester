package logs

import (
	"fmt"
	"testing"
)

func TestnewConsole_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cw := newConsole()

	if cw == nil {
		t.Error("newConsole() returned nil")
	}

	if cw.lg == nil {
		t.Error("cw.lg is nil")
	}

	if cw.Level != LevelDebug {
		t.Errorf("cw.Level is not LevelDebug, it's %d", cw.Level)
	}

	if !cw.Colorful {
		t.Error("cw.Colorful is not true")
	}

	if cw.formatter == nil {
		t.Error("cw.formatter is nil")
	}
}
