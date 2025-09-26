package htesting

import (
	"fmt"
	"regexp"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRun_4(t *testing.T) {
	testCases := []struct {
		name     string
		re       *regexp.Regexp
		testName string
		want     bool
	}{
		{
			name:     "match",
			re:       regexp.MustCompile("test"),
			testName: "test",
			want:     true,
		},
		{
			name:     "no match",
			re:       regexp.MustCompile("test"),
			testName: "notest",
			want:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &qt.C{TB: t}
			r := &PinnedRunner{c: c, re: tc.re}
			got := r.Run(tc.testName, func(c *qt.C) {})
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
