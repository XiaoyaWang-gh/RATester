package loggers

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestPrintTimerIfDelayed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	l := &logAdapter{}
	// CONTEXT_1
	// CONTEXT_2
	start := time.Now()
	// CONTEXT_3
	// CONTEXT_4
	// CONTEXT_5
	name := "test"
	// TEST CASE
	l.PrintTimerIfDelayed(start, name)
	// ASSERT
	assert.Equal(t, 0, len(l.errors.String()))
}
