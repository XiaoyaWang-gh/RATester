package logs

import (
	"fmt"
	"testing"
)

func TestNewConsole_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cw := newConsole()

	if cw == nil {
		t.Errorf("newConsole() = %v, want not nil", cw)
	}

	if cw.lg == nil {
		t.Errorf("newConsole().lg = %v, want not nil", cw.lg)
	}

	if cw.Level != LevelDebug {
		t.Errorf("newConsole().Level = %v, want %v", cw.Level, LevelDebug)
	}

	if !cw.Colorful {
		t.Errorf("newConsole().Colorful = %v, want true", cw.Colorful)
	}
}
