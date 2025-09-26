package gin

import (
	"fmt"
	"testing"
)

func TestSetMode_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"DebugMode", DebugMode, DebugMode},
		{"ReleaseMode", "release", ReleaseMode},
		{"TestMode", "test", TestMode},
		{"UnknownMode", "unknown", "gin mode unknown: unknown (available mode: debug release test)"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetMode(test.input)
			if modeName.Load() != test.expected {
				t.Errorf("expected %s, got %s", test.expected, modeName.Load())
			}
		})
	}
}
