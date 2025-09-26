package testenv

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestSkipFlakyNet_1(t *testing.T) {
	testCases := []struct {
		name     string
		env      string
		expected string
	}{
		{
			name:     "Test case 1",
			env:      "GO_BUILDER_FLAKY_NET=true",
			expected: "skipping test on builder known to have frequent network failures",
		},
		{
			name:     "Test case 2",
			env:      "GO_BUILDER_FLAKY_NET=false",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			os.Setenv("GO_BUILDER_FLAKY_NET", tc.env)
			defer os.Unsetenv("GO_BUILDER_FLAKY_NET")

			buf := &bytes.Buffer{}
			log.SetOutput(buf)
			defer log.SetOutput(os.Stderr)

			SkipFlakyNet(t)

			if tc.expected != "" && !strings.Contains(buf.String(), tc.expected) {
				t.Errorf("Expected log output to contain '%s', but got '%s'", tc.expected, buf.String())
			}
		})
	}
}
