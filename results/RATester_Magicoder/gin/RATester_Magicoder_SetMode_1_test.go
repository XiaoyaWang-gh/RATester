package gin

import (
	"fmt"
	"testing"
)

func TestSetMode_1(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"TestDebugMode", DebugMode, DebugMode},
		{"TestReleaseMode", ReleaseMode, ReleaseMode},
		{"TestTestMode", TestMode, TestMode},
		{"TestEmptyString", "", DebugMode},
		{"TestInvalidMode", "invalid", "gin mode unknown: invalid (available mode: debug release test)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetMode(tt.value)
			if modeName.Load() != tt.expected {
				t.Errorf("SetMode(%s) = %s, want %s", tt.value, modeName.Load(), tt.expected)
			}
		})
	}
}
