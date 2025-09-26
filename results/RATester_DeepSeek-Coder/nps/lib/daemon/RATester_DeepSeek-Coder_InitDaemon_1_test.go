package daemon

import (
	"fmt"
	"os"
	"testing"
)

func TestInitDaemon_1(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "TestStart",
			args:     []string{"start"},
			expected: "start",
		},
		{
			name:     "TestStop",
			args:     []string{"stop"},
			expected: "stop",
		},
		{
			name:     "TestRestart",
			args:     []string{"restart"},
			expected: "restart",
		},
		{
			name:     "TestStatus",
			args:     []string{"status"},
			expected: "status",
		},
		{
			name:     "TestReload",
			args:     []string{"reload"},
			expected: "reload",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Args = tc.args
			InitDaemon("f", "runPath", "pidPath")
			got := os.Args[1]
			if got != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, got)
			}
		})
	}
}
