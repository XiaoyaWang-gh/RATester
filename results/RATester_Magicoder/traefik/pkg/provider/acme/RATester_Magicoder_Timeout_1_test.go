package acme

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	challenge := &ChallengeHTTP{}
	timeout, interval := challenge.Timeout()

	if timeout != 60*time.Second {
		t.Errorf("Expected timeout to be 60 seconds, but got %v", timeout)
	}

	if interval != 5*time.Second {
		t.Errorf("Expected interval to be 5 seconds, but got %v", interval)
	}
}
