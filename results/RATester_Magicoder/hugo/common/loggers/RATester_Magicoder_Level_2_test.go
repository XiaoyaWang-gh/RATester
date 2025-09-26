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

	l := &logAdapter{
		level: logg.LevelInfo,
	}

	if l.Level() != logg.LevelInfo {
		t.Errorf("Expected level to be %v, but got %v", logg.LevelInfo, l.Level())
	}
}
