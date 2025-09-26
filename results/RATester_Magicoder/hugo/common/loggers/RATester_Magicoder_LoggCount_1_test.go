package loggers

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestLoggCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{
		logCounters: &logLevelCounter{
			counters: map[logg.Level]int{
				logg.LevelDebug: 1,
				logg.LevelInfo:  2,
				logg.LevelWarn:  3,
				logg.LevelError: 4,
			},
		},
	}

	tests := []struct {
		level logg.Level
		count int
	}{
		{logg.LevelDebug, 1},
		{logg.LevelInfo, 2},
		{logg.LevelWarn, 3},
		{logg.LevelError, 4},
		{logg.LevelTrace, 0}, // Test for unknown level
	}

	for _, test := range tests {
		got := l.LoggCount(test.level)
		if got != test.count {
			t.Errorf("LoggCount(%v) = %v, want %v", test.level, got, test.count)
		}
	}
}
