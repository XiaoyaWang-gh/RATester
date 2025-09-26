package tcp

import (
	"fmt"
	"testing"
	"time"
)

func TestTerminationDelay_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dialer := dialerWrapper{
		terminationDelay: 10 * time.Second,
	}

	expected := 10 * time.Second
	actual := dialer.TerminationDelay()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
