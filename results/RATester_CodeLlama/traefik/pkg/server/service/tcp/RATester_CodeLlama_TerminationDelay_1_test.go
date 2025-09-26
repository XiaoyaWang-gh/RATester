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

	d := dialerWrapper{
		terminationDelay: 10 * time.Second,
	}
	if got := d.TerminationDelay(); got != 10*time.Second {
		t.Errorf("TerminationDelay() = %v, want %v", got, 10*time.Second)
	}
}
