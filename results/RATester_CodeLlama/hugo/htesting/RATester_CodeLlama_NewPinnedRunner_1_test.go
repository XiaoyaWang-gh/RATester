package htesting

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewPinnedRunner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	pinnedTestRe := ".*"
	pinnedTestRe = strings.ReplaceAll(pinnedTestRe, "_", " ")
	re := regexp.MustCompile("(?i)" + pinnedTestRe)
	pinnedRunner := NewPinnedRunner(t, pinnedTestRe)
	assert.Equal(t, re, pinnedRunner.re)
}
