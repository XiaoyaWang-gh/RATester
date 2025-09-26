package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestLevel_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	l.level = logg.LevelDebug
	if l.Level() != logg.LevelDebug {
		t.Errorf("Level() = %v, want %v", l.Level(), logg.LevelDebug)
	}
}
