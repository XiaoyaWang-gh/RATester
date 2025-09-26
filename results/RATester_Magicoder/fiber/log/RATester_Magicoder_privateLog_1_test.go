package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestprivateLog_1(t *testing.T) {
	logger := &defaultLogger{
		stdlog: log.New(os.Stdout, "", log.LstdFlags),
		level:  LevelInfo,
		depth:  1,
	}

	tests := []struct {
		name     string
		level    Level
		fmtArgs  []any
		expected string
	}{
		{
			name:     "Debug",
			level:    LevelDebug,
			fmtArgs:  []any{"debug message"},
			expected: "DEBUG: debug message",
		},
		{
			name:     "Info",
			level:    LevelInfo,
			fmtArgs:  []any{"info message"},
			expected: "INFO: info message",
		},
		{
			name:     "Panic",
			level:    LevelPanic,
			fmtArgs:  []any{"panic message"},
			expected: "PANIC: panic message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			logger.privateLog(tt.level, tt.fmtArgs)

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = rescueStdout

			if !strings.Contains(string(out), tt.expected) {
				t.Errorf("Expected: %s, Got: %s", tt.expected, string(out))
			}
		})
	}
}
