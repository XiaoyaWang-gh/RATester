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
			name:     "start",
			args:     []string{"start"},
			expected: "start",
		},
		{
			name:     "stop",
			args:     []string{"stop"},
			expected: "stop",
		},
		{
			name:     "restart",
			args:     []string{"restart"},
			expected: "restart",
		},
		{
			name:     "status",
			args:     []string{"status"},
			expected: "status",
		},
		{
			name:     "reload",
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
			// assert.Equal(t, tc.expected, actual)
		})
	}
}
