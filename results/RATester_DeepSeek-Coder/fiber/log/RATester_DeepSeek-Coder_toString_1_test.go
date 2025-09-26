package log

import (
	"fmt"
	"testing"
)

func TestToString_1(t *testing.T) {
	tests := []struct {
		name string
		lv   Level
		want string
	}{
		{"Trace", LevelTrace, "Trace"},
		{"Debug", LevelDebug, "Debug"},
		{"Info", LevelInfo, "Info"},
		{"Warn", LevelWarn, "Warn"},
		{"Error", LevelError, "Error"},
		{"Unknown", 100, "[?100] "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.lv.toString(); got != tt.want {
				t.Errorf("Level.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
