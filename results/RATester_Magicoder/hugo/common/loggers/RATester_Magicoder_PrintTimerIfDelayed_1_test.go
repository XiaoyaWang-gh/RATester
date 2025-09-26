package loggers

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintTimerIfDelayed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	start := time.Now()
	name := "test"

	l := &logAdapter{}
	l.PrintTimerIfDelayed(start, name)

	elapsed := time.Since(start)
	milli := int(1000 * elapsed.Seconds())

	if milli < 500 {
		return
	}

	expected := fmt.Sprintf("%s in %v ms", name, milli)

	if l.sprint() != expected {
		t.Errorf("Expected %s, but got %s", expected, l.sprint())
	}
}
